package tests

import (
	"testing"

	"github.com/defUserName-404/coding-challenges/challenges/004-webserver/src"
)

func TestPlaceholder(t *testing.T) {
	if webserver.Placeholder() == "" {
		t.Fatal("Placeholder() returned empty string")
	}
}
