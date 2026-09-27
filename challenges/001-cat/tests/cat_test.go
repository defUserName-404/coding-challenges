package tests

import (
	"testing"

	"github.com/defUserName-404/coding-challenges/challenges/001-cat/src"
)

func TestPlaceholder(t *testing.T) {
	if cat.Placeholder() == "" {
		t.Fatal("Placeholder() returned empty string")
	}
}
