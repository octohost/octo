// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
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
)

func init() {
	RootCmd.AddCommand(configCmd)
	configCmd.PersistentFlags().StringVarP(&Container, "container", "c", "", "Docker Container")
}
