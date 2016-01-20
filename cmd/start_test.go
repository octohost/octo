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
