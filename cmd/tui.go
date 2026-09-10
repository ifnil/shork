package cmd

import (
	"github.com/ifnil/shork/internal/groups"
	"github.com/ifnil/shork/internal/host"
	"github.com/ifnil/shork/internal/runner"
	"github.com/ifnil/shork/internal/sshx"
	"github.com/ifnil/shork/tui"
	"github.com/spf13/cobra"
)

var tuiCmd = &cobra.Command{
	Use:   "tui",
	Short: "A brief description of your command",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		agc, err := sshx.GetAgentClient()
		if err != nil {
			agc = nil
		}

		kr := sshx.NewKeyring(agc)
		gm := groups.BuildGroupMap()
		hm := host.NewHostMap()
		if err := hm.LoadSSHConfig("~/.ssh/config"); err != nil {
			return err
		}

		rr := runner.NewRunner(kr, hm, gm)
		if err := tui.Run(cmd.Context(), rr); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(tuiCmd)
}
