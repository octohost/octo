// +build linux darwin freebsd

package cmd

import (
	"testing"
)

func TestConsulConfigPath(t *testing.T) {
	Container = "testing"
	ConfigKey = "octo"
	config := ConfigEnv{Container: Container, Key: ConfigKey}
	fullPath := config.Path()
	if fullPath != "octohost/testing/OCTO" {
		t.Errorf("The ConfigPath() was incorrect: %s", fullPath)
	}
}

func TestConsulSetEnvVar(t *testing.T) {
	Container = "testing"
	ConfigKey = "octo"
	ConfigValue = "This is the value for the octo key."
	config := ConfigEnv{Container: Container, Key: ConfigKey, Value: ConfigValue}
	fullPath := config.Path()
	if ConfigSet(fullPath, ConfigValue) {
		t.Logf("Set the key %s.", fullPath)
		value := ConfigGet(fullPath)
		if value != "This is the value for the octo key." {
			t.Errorf("The env var is NOT correct.")
		}
	}
}
