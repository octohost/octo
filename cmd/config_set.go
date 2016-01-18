// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
	config := ConfigEnv{Container: Container, Key: ConfigKey, Value: ConfigValue}
	if config.Set() {
		Log(fmt.Sprintf("set key='%s'", ConfigKey), "info")
	} else {
		fmt.Printf("Error: config set key='%s'\n", ConfigKey)
	}
}

func checkConfigSetFlags() {
	Log("Checking cli flags.", "debug")
	if Container == "" {
		fmt.Println("A container is required: -c")
		os.Exit(1)
	}
	if ConfigKey == "" {
		fmt.Println("A key is required: --key")
		os.Exit(1)
	}
	if ConfigValue == "" {
		fmt.Println("A value is required: --value")
		os.Exit(1)
	}
	Log("Required cli flags are present.", "debug")
}

func init() {
	configCmd.AddCommand(configSetCmd)
}
