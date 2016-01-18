// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
	config := ConfigEnv{Container: Container, Key: ConfigKey}
	value := config.Get()
	if value != "" {
		fmt.Printf("%s\n", value)
	}
}

func checkConfigGetFlags() {
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
	configCmd.AddCommand(configGetCmd)
}

// Get returns the value of the key passed.
func (c *ConfigEnv) Get() string {
	consul, err := ConsulSetup()
	if err != nil {
		Log("Fatal Consul setup problem.", "info")
	}
	value := ConsulGet(consul, c.Path())
	return value
}
