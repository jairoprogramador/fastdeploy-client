package cmd

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jairoprogramador/fastdeploy/internal/application"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/docker"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/executor"
	infraProject "github.com/jairoprogramador/fastdeploy/internal/infrastructure/project"
	"github.com/mattn/go-isatty"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "0.1.0"
	withTty bool
)

var rootCmd = &cobra.Command{
	Use:   "fd [step] [environment]",
	Short: "fastdeploy is a tool for managing your deployment workflow.",
	Long:  `A flexible command-line tool to handle authentication and execution of various deployment steps.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {

			if cmd.HasSubCommands() && cmd.CalledAs() == "fd" {
				return nil
			}
			return errors.New("a step argument is required")
		}
		if len(args) > 2 {
			return errors.New("a maximum of two arguments are allowed: a step and an optional environment")
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// Silence usage on error to avoid printing help on execution errors.
		cmd.SilenceUsage = true

		if len(args) == 0 {
			return cmd.Help()
		}

		// --- Pre-flight checks ---
		workDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get working directory: %w", err)
		}

		projRepo := infraProject.NewYAMLProjectRepository(workDir)
		exists, err := projRepo.Exists()
		if err != nil {
			return fmt.Errorf("could not check for project initialization: %w", err)
		}
		if !exists {
			return errors.New("project not initialized. Please run 'fd init' first")
		}

		isTerminal := isatty.IsTerminal(os.Stdout.Fd())

		// --- Composition Root ---
		projectRepo := infraProject.NewYAMLProjectRepository(workDir)
		cmdExecutor := executor.NewShellCommandExecutor()
		dockerService := docker.NewManagerDockerService(cmdExecutor)
		runStepService := application.NewRunStepService(isTerminal, workDir, getFastdeployHome(), projectRepo, dockerService)
		// ---

		step := args[0]
		env := ""
		if len(args) > 1 {
			env = args[1]
		}

		return runStepService.ExecuteStep(cmd.Context(), step, env, withTty)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&withTty, "with-tty", false, "Enable pseudo-TTY allocation.")
	rootCmd.Version = fmt.Sprintf("v%s\n", version)
	rootCmd.SetVersionTemplate(`{{.Version}}`)

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)
}

func getFastdeployHome() string {
	viper.SetEnvPrefix("FASTDEPLOY")
	viper.AutomaticEnv()

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error al obtener el directorio home:", err)
		os.Exit(1)
	}

	defaultHome := filepath.Join(userHomeDir, ".fastdeploy")
	fastdeployHome := viper.GetString("HOME")
	if fastdeployHome == "" {
		fastdeployHome = defaultHome
	}
	return fastdeployHome
}
