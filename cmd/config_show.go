// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
	config := ConfigEnv{Container: Container}
	configs := config.Variables()
	for _, c := range configs {
		c.Show()
	}
}

func checkConfigShowFlags() {
	Log("Checking cli flags.", "debug")
	if Container == "" {
		fmt.Println("A container is required: -c")
		os.Exit(1)
	}
	Log("Required cli flags are present.", "debug")
}

func init() {
	configCmd.AddCommand(configShowCmd)
}
