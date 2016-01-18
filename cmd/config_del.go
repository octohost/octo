// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
	config := ConfigEnv{Container: Container, Key: ConfigKey}
	value := config.Del()
	if value {
		Log(fmt.Sprintf("Deleted %s", config.Path()), "info")
	}
}

func checkConfigDelFlags() {
	Log("Checking cli flags.", "debug")
	if Container == "" {
		fmt.Println("A container is required: -c")
		os.Exit(1)
	}
	if ConfigKey == "" {
		fmt.Println("A key is required: --key")
		os.Exit(1)
	}
	Log("Required cli flags are present.", "debug")
}

func init() {
	configCmd.AddCommand(configDelCmd)
}
