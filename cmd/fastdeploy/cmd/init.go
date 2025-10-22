package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/jairoprogramador/fastdeploy/internal/application"
	infra "github.com/jairoprogramador/fastdeploy/internal/infrastructure/project"
	"github.com/spf13/cobra"
)

var nonInteractive bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new project (creates fdconfig.yaml).",
	RunE: func(cmd *cobra.Command, args []string) error {
		workDir, err := os.Getwd()
		if err != nil {
			return fmt.Errorf("could not get working directory: %w", err)
		}

		repository := infra.NewYAMLProjectRepository(workDir)
		inputService := infra.NewSurveyUserInputService()
		initService := application.NewInitService(filepath.Base(workDir), repository, inputService)

		err = initService.InitializeProject(cmd.Context(), !nonInteractive)
		if err != nil {
			if err == application.ErrProjectAlreadyExists {
				fmt.Println(application.MessageProjectAlreadyExists)
				return nil
			}
			return fmt.Errorf("initialization failed: %w", err)
		}

		fmt.Println("✅ Project initialized successfully. 'fdconfig.yaml' created.")
		return nil
	},
}

func init() {
	initCmd.Flags().BoolVarP(&nonInteractive, "yes", "y", false, "Use default values without prompting for input.")
}
