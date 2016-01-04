// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var configDelCmd = &cobra.Command{
	Use:     "del",
	Aliases: []string{"rm"},
	Short:   "Delete configuration for Docker container from Consul.",
	PreRun: func(cmd *cobra.Command, args []string) {
		checkConfigDelFlags()
		LoadConfig()
	},
	Long: `Delete configuration for Docker container from Consul.`,
	Run:  startConfigDel,
}

func startConfigDel(cmd *cobra.Command, args []string) {
	fmt.Println("Del")
}

func checkConfigDelFlags() {
	Log("Checking flags", "info")
}

func init() {
	configCmd.AddCommand(configDelCmd)
}
