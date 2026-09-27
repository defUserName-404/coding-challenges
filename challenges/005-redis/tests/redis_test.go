package tests

import (
	"testing"

	"github.com/defUserName-404/coding-challenges/challenges/005-redis/src"
)

func TestPlaceholder(t *testing.T) {
	if redis.Placeholder() == "" {
		t.Fatal("Placeholder() returned empty string")
	}
}
