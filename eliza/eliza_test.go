package eliza

import (
	"regexp"
	"testing"
)

// TestCompile tests the regex compilation
func TestCompile(t *testing.T) {
	for pattern := range Psychobabble {
		_, err := regexp.Compile(pattern)
		if err != nil {
			t.Errorf("Cannot compile %s", pattern)
		}
	}
}
