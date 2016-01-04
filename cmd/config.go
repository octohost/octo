// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"strings"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Get, set, delete and show configuration in Consul for Docker containers.",
	Long:  `Get, set, delete and show configuration in Consul for Docker containers.`,
	Run:   startConfig,
}

func startConfig(cmd *cobra.Command, args []string) {
	fmt.Println("octo config -h")
}

var (
	// Container is the Docker container we are loading config for.
	Container string

	// ConfigKey is the key for the ENV variable.
	ConfigKey string

	// ConfigValue is the value for the ENV variable.
	ConfigValue string
)

func init() {
	RootCmd.AddCommand(configCmd)
	configCmd.PersistentFlags().StringVarP(&Container, "container", "c", "", "Docker Container")
	configCmd.PersistentFlags().StringVarP(&ConfigKey, "key", "", "", "Key for environmental variable.")
	configCmd.PersistentFlags().StringVarP(&ConfigValue, "value", "", "", "Value for environmental variable.")
}

// ConfigPath returns the entire Consul path for a Consul config variable.
func ConfigPath() string {
	prefix := ""
	if prefix = viper.GetString("prefix"); prefix == "" {
		prefix = ConsulPrefix
	}
	fullPath := fmt.Sprintf("%s/%s/%s", strings.TrimPrefix(prefix, "/"), Container, strings.ToUpper(ConfigKey))
	return fullPath
}
