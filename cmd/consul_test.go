package cmd

import (
	"testing"
)

func TestCleanupToken(t *testing.T) {
	token := "this-is-the-token"
	cleanedToken := cleanupToken(token)
	if cleanedToken != "this" {
		t.Error("It didn't cleanup the token right.")
	}
}

func TestCleanupAnonymousToken(t *testing.T) {
	token := "anonymous"
	cleanedToken := cleanupToken(token)
	if cleanedToken != token {
		t.Error("It didn't cleanup the token right.")
	}
}
