// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/samalba/dockerclient"
	"github.com/spf13/viper"
	"strings"
)

const (
	buildOrg = "octoprod"
)

// Image is the struct for a Docker image.
type Image struct {
	Name     string
	BuildOrg string
	Tag      string
}

// GetBuildOrg returns the configured or default build_org.
func GetBuildOrg() string {
	org := ""
	if org = viper.GetString("build_org"); org == "" {
		org = buildOrg
	}
	return org
}

// GetTag gets the tag or assigns "latest"
func GetTag(passedName string) (string, string) {
	name := ""
	tag := ""
	if strings.Contains(passedName, ":") {
		imageAndTag := strings.Split(passedName, ":")
		name = imageAndTag[0]
		tag = imageAndTag[1]
	} else {
		name = passedName
		tag = "latest"
	}
	return name, tag
}

// GetImage creates an Image struct from the string we get from the CLI.
func GetImage(containerName string) Image {
	name := ""
	tag := ""
	if strings.Contains(containerName, "/") {
		parts := strings.Split(containerName, "/")
		name, tag = GetTag(parts[1])
		image := Image{Name: name, Tag: tag, BuildOrg: parts[0]}
		return image
	}
	buildOrg := GetBuildOrg()
	name, tag = GetTag(containerName)
	image := Image{Name: name, Tag: tag, BuildOrg: buildOrg}
	return image
}

// PrintName returns the full name of the Image.
func (i *Image) PrintName() string {
	name := fmt.Sprintf("%s/%s:%s", i.BuildOrg, i.Name, i.Tag)
	return name
}

// Start turns a docker image into a container.
func (i *Image) Start() {
	containerName := i.PrintName()

	// Make a connection to Docker.
	docker := DockerConnect()

	// Inspect the image.
	image := InspectImage(docker, containerName)

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
