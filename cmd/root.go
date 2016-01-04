// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// RootCmd is the default Cobra struct that starts it all off.
// https://github.com/spf13/cobra
var RootCmd = &cobra.Command{
	Use:   "octo",
	Short: "CLI command for octohost.",
	Long:  `CLI command for octohost. This is how you interact with all octohost components.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("`octo -h` for help information.")
		fmt.Println("`octo -v` for version information.")
	},
}

const (
	// ConsulPrefix is the default location where we will be storing all configuration.
	// It can be overwritten with the "prefix" value in the config.yaml file.
	ConsulPrefix = "octohost"
)

var (
	// Direction adds information about which command is running to the logs.
	Direction string

	// Verbose logs all output to stdout.
	Verbose bool
)

func init() {
	Direction = SetDirection()
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "", false, "log output to stdout")
}
