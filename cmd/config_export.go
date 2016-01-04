// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
	fmt.Println("Export")
}

func checkConfigExportFlags() {
	Log("Checking flags", "info")
}

func init() {
	configCmd.AddCommand(configExportCmd)
}
