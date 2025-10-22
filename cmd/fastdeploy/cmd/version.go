package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	structColor    = color.New(color.FgWhite)
	commandColor   = color.New(color.FgGreen, color.Bold)
	versionColor   = color.New(color.FgHiGreen)
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of the CLI.",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println()
		structColor.Print("                  . ")
		structColor.Print("4242")
		structColor.Println(" .")
		structColor.Println("                 1        1")
		structColor.Println("              ==1==========1==")
		structColor.Println("                1  @    @  1")
		structColor.Println("                 1        1")
		structColor.Println("                   1    1")
		structColor.Print("              .")
		structColor.Print("1337")
		structColor.Print("      ")
		structColor.Print("1337")
		structColor.Println(".")
		structColor.Println("            @                  @")
		structColor.Println("          1                      1")
		structColor.Print("         0  ")
		structColor.Print("100000000")
		structColor.Print("📸")
		structColor.Print("100000000")
		structColor.Println("  0         ")
		structColor.Println("        1   1                  1   1")
		structColor.Print("        0   0  ")
		commandColor.Print("$ fd init")
		structColor.Println("       0   0")
		structColor.Print("        0   0  ")
		commandColor.Print("$ fd test")
		structColor.Println("       0   0")
		structColor.Print("        0   0  ")
		commandColor.Print("$ fd deploy")
		structColor.Println("     0   0")

		structColor.Println("          @ @__________________@ 0")
		structColor.Print("            ")
		structColor.Print("110111011-0xDEADBEEF")

		fmt.Println()
		fmt.Println()

		versionStr := fmt.Sprintf("       FastDeploy CLI version: v%s", version)
		versionColor.Println(versionStr)
		fmt.Println()
	},
}
