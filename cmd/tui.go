package cmd

import (
	"github.com/ifnil/shork/tui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "A brief description of your command",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if err := tui.Run(cmd.Context()); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
