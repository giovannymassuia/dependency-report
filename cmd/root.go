package cmd

import (
	"fmt"
	cliVerison "github.com/giovannymassuia/dependency-report/version"
	"os"

	"github.com/spf13/cobra"

	cc "github.com/ivanpirog/coloredcobra"
)

var rootCmd = &cobra.Command{
	Use:   "dependency-report",
	Short: "A tool to generate dependency reports",
	Long:  `A tool to generate dependency reports from various sources.`,
	Run: func(cmd *cobra.Command, args []string) {

		// version flag
		if version, _ := cmd.Flags().GetBool("version"); version {
			// get version from last git tag
			fmt.Printf("dependency-report current version %s\n", cliVerison.CurrVersion)
			return
		}

		cmd.Help()
	},
}

func Execute() {
	cc.Init(&cc.Config{
		RootCmd:  rootCmd,
		Headings: cc.HiCyan + cc.Bold + cc.Underline,
		Commands: cc.HiYellow + cc.Bold,
		Example:  cc.Italic,
		ExecName: cc.Bold,
		Flags:    cc.Bold,
	})
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// version flag
	rootCmd.Flags().BoolP("version", "v", false, "Print version information and quit")
}
