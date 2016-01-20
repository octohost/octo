// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"strings"
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

const (
	buildOrg = "octoprod"
)

func startStart(cmd *cobra.Command, args []string) {
	//image := GetImage(Container)
}

// Image is the struct for a Docker image.
type Image struct {
	Name     string
	BuildOrg string
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

// GetImage creates an Image struct from the string we get from the CLI.
func GetImage(containerName string) Image {
	if strings.Contains(containerName, "/") {
		parts := strings.Split(containerName, "/")
		image := Image{Name: parts[1], BuildOrg: parts[0]}
		return image
	}
	buildOrg := GetBuildOrg()
	image := Image{Name: containerName, BuildOrg: buildOrg}
	return image
}
