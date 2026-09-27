package tests

import (
	"testing"

	"github.com/defUserName-404/coding-challenges/challenges/002-wc/src"
)

func TestPlaceholder(t *testing.T) {
	if wc.Placeholder() == "" {
		t.Fatal("Placeholder() returned empty string")
	}
}
