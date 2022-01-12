package cmd

import (
	"github.com/ImperiumProject/imperium/cmd/strategies"
	"github.com/ImperiumProject/imperium/cmd/visualizer"
	"github.com/ImperiumProject/imperium/config"
	"github.com/spf13/cobra"
)

// RootCmd returns the root cobra command of the scheduler tool
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "imperium",
		Short: "Tool to test/understand consensus algorithms",
	}
	cmd.CompletionOptions.DisableDefaultCmd = true
	cmd.PersistentFlags().StringVarP(&config.ConfigPath, "config", "c", "config.json", "Config file path")
	cmd.AddCommand(visualizer.VisualizerCmd())
	cmd.AddCommand(strategies.StrategiesCmd())
	return cmd
}
