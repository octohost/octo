// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var configExportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export configuration for Docker container from Consul.",
	PreRun: func(cmd *cobra.Command, args []string) {
		checkConfigExportFlags()
		LoadConfig()
	},
	Long: `Export configuration for Docker container from Consul.`,
	Run:  startConfigExport,
}

func startConfigExport(cmd *cobra.Command, args []string) {
	config := ConfigEnv{Container: Container}
	configs := config.Variables()
	for _, c := range configs {
		c.Export()
	}
}

func checkConfigExportFlags() {
	Log("Checking cli flags.", "debug")
	if Container == "" {
		fmt.Println("A container is required: -c")
		os.Exit(1)
	}
	SpaceCheck(Container, "container")
	Log("Required cli flags are present.", "debug")
}

func init() {
	configCmd.AddCommand(configExportCmd)
}
