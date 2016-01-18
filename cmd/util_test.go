package cmd

import (
	"testing"
)

func TestSetDirectionBinary(t *testing.T) {
	direction := SetDirection("bin/octo")
	if direction != "main" {
		t.Error("Wrong direction.")
	}
}

func TestSetDirectionWithHelp(t *testing.T) {
	direction := SetDirection("bin/octo -h")
	if direction != "main" {
		t.Error("Wrong direction.")
	}
}

func TestSetDirectionConfig(t *testing.T) {
	direction := SetDirection("bin/octo config")
	if direction != "config" {
		t.Error("Wrong direction.")
	}
}
