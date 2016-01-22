// +build linux darwin freebsd

package cmd

import (
	"fmt"
	"github.com/samalba/dockerclient"
	"github.com/spf13/viper"
	"os"
)

const (
	dockerSock = "unix:///var/run/docker.sock"
)

// DockerConnect connects to Docker and returns a client.
func DockerConnect() *dockerclient.DockerClient {
	sock := ""
	if sock = viper.GetString("docker_sock"); sock == "" {
		sock = dockerSock
	}
	docker, err := dockerclient.NewDockerClient(sock, nil)
	if err != nil {
		Log("There was a docker connection error", "info")
		os.Exit(1)
	}
	Log("Connected to Docker", "info")
	return docker
}

// InspectImage tries to inspect the image and pulls it if it's not found.
func InspectImage(d *dockerclient.DockerClient, image string) *dockerclient.ImageInfo {
	Log(fmt.Sprintf("Inspecting the image: '%s'", image), "info")
	imageD, err := d.InspectImage(image)
	if err != nil {
		Log(fmt.Sprintf("Image not found: '%s'", image), "info")
		errPull := PullImage(d, image)
		if errPull == nil {
			imageD = InspectImage(d, image)
		} else {
			Log(fmt.Sprintf("There is no image with that name: '%s'", image), "info")
			os.Exit(1)
		}
	}
	return imageD
}

// PullImage gets a docker image from the Docker Hub.
func PullImage(d *dockerclient.DockerClient, image string) error {
	Log(fmt.Sprintf("Pulling Image: '%s'", image), "info")
	err := d.PullImage(image, nil)
	return err
}

// ListImages returns all the images we are looking for.
func ListImages(d *dockerclient.DockerClient) ([]*dockerclient.Image, error) {
	images, err := d.ListImages(true)
	return images, err
}
