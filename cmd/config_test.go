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
	if config.Set() {
		t.Logf("Set the key %s.", config.Path())
		value := config.Get()
		if value != "This is the value for the octo key." {
			t.Errorf("The env var is NOT correct.")
		}
	}
}

func TestConsulDelEnvVar(t *testing.T) {
	Container = "testing"
	ConfigKey = "octo"
	ConfigValue = "This is the value for the octo key."
	config := ConfigEnv{Container: Container, Key: ConfigKey, Value: ConfigValue}
	if config.Set() {
		if config.Del() {
			value := config.Get()
			if value != "" {
				t.Error("Should not have been able to get it.")
			}
		}
	}
}
