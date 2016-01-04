// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var configShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Show configuration for Docker container from Consul.",
	PreRun: func(cmd *cobra.Command, args []string) {
		checkConfigShowFlags()
		LoadConfig()
	},
	Long: `Show configuration for Docker container from Consul.`,
	Run:  startConfigShow,
}

func startConfigShow(cmd *cobra.Command, args []string) {
	fmt.Println("Show")
}

func checkConfigShowFlags() {
	Log("Checking flags", "info")
}

func init() {
	configCmd.AddCommand(configShowCmd)
}
