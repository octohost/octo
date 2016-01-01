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

}

func checkConfigFlags() {
	Log("Checking flags", "info")
}

func init() {
	RootCmd.AddCommand(configCmd)
}
