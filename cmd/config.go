// +build linux darwin freebsd

package cmd

import (
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get and set configuration in Consul for Docker containers.",
	PreRun: func(cmd *cobra.Command, args []string) {
		checkConfigFlags()
		LoadConfig()
	},
	Long: `Get and set configuration in Consul for Docker containers.`,
	Run:  startConfig,
}

func startConfig(cmd *cobra.Command, args []string) {
	consul, err := ConsulSetup()
	if err != nil {
		Log("Fatal Consul setup problem.", "info")
	}
	if ConsulSet(consul, "testing", "the-value") {
		Log("Set the value at testing", "info")
	}
	value, err := ConsulGet(consul, "testing")
	if value == "" {
		Log("Nothing at that key.", "info")
	}
	if ConsulDel(consul, "testing") {
		Log("Removed the value at testing", "info")
	}
}

func checkConfigFlags() {
	Log("Checking flags", "info")
}

func init() {
	RootCmd.AddCommand(configCmd)
}
