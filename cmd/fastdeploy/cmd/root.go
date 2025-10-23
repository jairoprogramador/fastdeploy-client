package cmd

import (
	"errors"
	"fmt"
	"os"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/factories"
	"github.com/spf13/cobra"
)

var (
	version = "0.1.0"
	withTty bool
)

var rootCmd = &cobra.Command{
	Use:   "fastdeploy",
	Short: "fastdeploy is a CLI tool for managing and deploying projects",
	Long:  `fastdeploy is a powerful and flexible CLI tool designed to streamline your development and deployment workflows.`,
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
		cmd.SilenceUsage = true

		if len(args) == 0 {
			return cmd.Help()
		}

		order := args[0]
		env := ""
		if len(args) > 1 {
			env = args[1]
		}

		factory := factories.NewServiceFactory()
		logFile, err := factory.BuildFileLogRepository().CreateFile()
		if err != nil {
			return err
		}
		defer logFile.Close()

		orderService, err := factory.BuildOrderService(logFile)
		if err != nil {
			return err
		}
		return orderService.ExecuteOrder(cmd.Context(), order, env, withTty)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&withTty, "with-tty", false, "Enable pseudo-TTY allocation.")
	rootCmd.Version = fmt.Sprintf("v%s\n", version)
	rootCmd.SetVersionTemplate(`{{.Version}}`)

	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(logCmd)
}
