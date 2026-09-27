package tests

import (
	"testing"

	"github.com/defUserName-404/coding-challenges/challenges/003-shell/src"
)

func TestPlaceholder(t *testing.T) {
	if shell.Placeholder() == "" {
		t.Fatal("Placeholder() returned empty string")
	}
}
