package cmd

import (
	"os"

	"github.com/spf13/cobra"

	cc "github.com/ivanpirog/coloredcobra"
)

var rootCmd = &cobra.Command{
	Use:   "dependency-report",
	Short: "A tool to generate dependency reports",
	Long:  `A tool to generate dependency reports from various sources.`,
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
}
