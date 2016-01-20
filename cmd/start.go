// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a Docker conatiner.",
	Long:  `Start a Docker conatiner.`,
	Run:   startStart,
}

const (
	buildOrg = "octoprod"
)

func startStart(cmd *cobra.Command, args []string) {
	fmt.Println("octo start -h")
}

var ()

func init() {
	RootCmd.AddCommand(startCmd)
}

// GetBuildOrg returns the configured or default build_org.
func GetBuildOrg() string {
	org := ""
	if org = viper.GetString("build_org"); org == "" {
		org = buildOrg
	}
	return org
}
