package cmd

import (
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/factories"
	"github.com/spf13/cobra"
)

var nonInteractive bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a new project (creates fdconfig.yaml).",
	RunE: func(cmd *cobra.Command, args []string) error {
		factory := factories.NewServiceFactory()

		logFile, err := factory.BuildFileLogRepository().CreateFile()
		if err != nil {
			return err
		}
		defer logFile.Close()

		initService, err := factory.BuildInitService(logFile)
		if err != nil {
			return err
		}

		if err = initService.InitializeProject(cmd.Context(), !nonInteractive);err != nil {
			return err
		}
		return nil
	},
}

func init() {
	initCmd.Flags().BoolVarP(&nonInteractive, "yes", "y", false, "Use default values without prompting for input.")
}
