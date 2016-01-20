package cmd

import (
	"testing"
)

// This will always get the constant because it doesn't LoadConfig().
func TestGetBuildOrg(t *testing.T) {
	org := GetBuildOrg()
	if org != "octoprod" {
		t.Errorf("Wrong default build org: %s", org)
	}
}

func TestCreateStandardImage(t *testing.T) {
	container := "darron"
	image := GetImage(container)
	if image.Name != "darron" {
		t.Errorf("The image.Name is incorrect: %s", image.Name)
	}
	if image.BuildOrg != "octoprod" {
		t.Errorf("The image.BuildOrg is incorrect: %s", image.BuildOrg)
	}
}

func TestCreateCustomImage(t *testing.T) {
	container := "octohost/darron"
	image := GetImage(container)
	if image.Name != "darron" {
		t.Errorf("The image.Name is incorrect: %s", image.Name)
	}
	if image.BuildOrg != "octohost" {
		t.Errorf("The image.BuildOrg is incorrect: %s", image.BuildOrg)
	}
}
