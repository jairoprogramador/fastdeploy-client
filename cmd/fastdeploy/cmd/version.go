package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	structColor    = color.New(color.FgHiGreen, color.Bold)
	versionColor   = color.New(color.FgWhite)
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of the CLI.",
	Run: func(cmd *cobra.Command, args []string) {

		structColor.Println(" _____         _   ____             _")
		structColor.Println("|  ___|_ _ ___| |_|  _ \\  ___ _ __ | | ___  _   _")
		structColor.Println("| |_ / _` / __| __| | | |/ _ \\ '_ \\| |/ _ \\| | | |")
		structColor.Println("|  _| (_| \\__ \\ |_| |_| |  __/ |_) | | (_) | |_| |")
		structColor.Println("|_|  \\__,_|___/\\__|____/ \\___| .__/|_|\\___/ \\__, |")
		structColor.Println("                             |_|            |___/")
		
		fmt.Println()

		versionStr := fmt.Sprintf("CLI FastDeploy Client: v%s", version)
		versionColor.Println(versionStr)
		fmt.Println()
	},
}
