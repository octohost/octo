// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a Docker conatiner.",
	PreRun: func(cmd *cobra.Command, args []string) {
		checkStartFlags()
		LoadConfig()
	},
	Long: `Start a Docker conatiner.`,
	Run:  startStart,
}

func startStart(cmd *cobra.Command, args []string) {
	image := GetImage(Container)
	image.Start()
}

func checkStartFlags() {
	Log("Checking cli flags.", "debug")
	if Container == "" {
		fmt.Println("A container is required: -c")
		os.Exit(1)
	}
	SpaceCheck(Container, "container")
	Log("Required cli flags are present.", "debug")
}

func init() {
	RootCmd.AddCommand(startCmd)
}
