// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/samalba/dockerclient"
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
	image := GetImage(Container)
	image.Start()
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

// PrintName returns the full name of the Image.
func (i *Image) PrintName() string {
	name := fmt.Sprintf("%s/%s", i.BuildOrg, i.Name)
	return name
}

// Start turns a docker image into a container.
func (i *Image) Start() {
	containerName := i.PrintName()

	// Make a connection to Docker.
	docker, _ := dockerclient.NewDockerClient("unix:///var/run/docker.sock", nil)

	// Inspect the image.
	image, err := docker.InspectImage(containerName)
	Log(fmt.Sprintf("Image: '%s' Cmd: '%s'", containerName, image.ContainerConfig.Cmd), "info")

	// Setup the container.
	containerConfig := &dockerclient.ContainerConfig{
		Image:        containerName,
		Cmd:          image.Config.Cmd,
		ExposedPorts: image.Config.ExposedPorts}

	// Create a containerID
	containerID, err := docker.CreateContainer(containerConfig, "", nil)
	if err != nil {
		Log(fmt.Sprintf("%s", err), "info")
	}

	// Setups the hostConfig
	hostConfig := &dockerclient.HostConfig{}

	// Actually start the container.
	err = docker.StartContainer(containerID, hostConfig)
}
