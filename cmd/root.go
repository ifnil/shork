package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/ifnil/shork/internal/groups"
	"github.com/ifnil/shork/internal/host"
	"github.com/ifnil/shork/internal/runner"
	"github.com/ifnil/shork/internal/sshx"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// FUN: lol, make host implement io.Reader/io.Writer

var cfgFile string

var (
	runGroup string
)

var rootCmd = &cobra.Command{
	Use:   "shork",
	Short: "A brief description of your application",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
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
		if runGroup != "" {
			if err := rr.RunGroupCmd(ctx, runGroup, "uname -r"); err != nil {
				return err
			}
		}

		return nil
	},
}

func Execute(ctx context.Context) {
	err := rootCmd.ExecuteContext(ctx)
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "")
	rootCmd.PersistentFlags().StringVarP(&runGroup, "group", "g", "all", "group")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("config")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
