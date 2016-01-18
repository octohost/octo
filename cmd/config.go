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

// ConfigEnv is the struct for an environmental variable for a container.
type ConfigEnv struct {
	Container string
	Key       string
	Value     string
}

// Prefix returns the Consul path for the Container.
func (c *ConfigEnv) Prefix() string {
	prefix := ""
	if prefix = viper.GetString("prefix"); prefix == "" {
		prefix = ConsulPrefix
	}
	containerPath := fmt.Sprintf("%s/%s", strings.TrimPrefix(prefix, "/"), c.Container)
	return containerPath
}

// Path returns the entire Consul path for a Consul config variable.
func (c *ConfigEnv) Path() string {
	fullPath := fmt.Sprintf("%s/%s", c.Prefix(), strings.ToUpper(c.Key))
	return fullPath
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

// Set sets a key to a value for a container.
func (c *ConfigEnv) Set() bool {
	consul, err := ConsulSetup()
	if err != nil {
		Log("Fatal Consul setup problem.", "info")
	}
	if ConsulSet(consul, c.Path(), c.Value) {
		Log(fmt.Sprintf("ConfigSet key='%s'", c.Path()), "info")
		return true
	}
	return false
}

// Del deletes a key from Consul.
func (c *ConfigEnv) Del() bool {
	consul, err := ConsulSetup()
	if err != nil {
		Log("Fatal Consul setup problem.", "info")
	}
	value := ConsulDel(consul, c.Path())
	return value
}
