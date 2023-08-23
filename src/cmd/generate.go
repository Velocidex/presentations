package cmd

import (
	"github.com/Velocidex/presentations/src/generator"
	"github.com/spf13/cobra"
)

var (
	moduleFilterRegex = ""
)

// Generates the HTML files for the site.
var (
	generateCmd = &cobra.Command{
		Use:   "generate [flags] output_directory",
		Short: "Generate the site",
		Args:  cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
		RunE: func(cmd *cobra.Command, args []string) error {
			return generator.GenerateSite(
				args[0], Verbose, moduleFilterRegex)
		},
	}
)

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(
		&moduleFilterRegex, "filter", "", "", "A regex to filter only some modules")
}
