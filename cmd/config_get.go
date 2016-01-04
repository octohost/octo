// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var configGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get configuration from Consul for Docker container.",
	PreRun: func(cmd *cobra.Command, args []string) {
		checkConfigGetFlags()
		LoadConfig()
	},
	Long: `Get configuration from Consul for Docker container.`,
	Run:  startConfigGet,
}

func startConfigGet(cmd *cobra.Command, args []string) {
	fmt.Println("Get")
}

func checkConfigGetFlags() {
	Log("Checking flags", "info")
}

func init() {
	configCmd.AddCommand(configGetCmd)
}
