// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var configSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set configuration in Consul for Docker container.",
	PreRun: func(cmd *cobra.Command, args []string) {
		checkConfigSetFlags()
		LoadConfig()
	},
	Long: `Set configuration in Consul for Docker container.`,
	Run:  startConfigSet,
}

func startConfigSet(cmd *cobra.Command, args []string) {
	fmt.Println("Set")
}

func checkConfigSetFlags() {
	Log("Checking flags", "info")
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
