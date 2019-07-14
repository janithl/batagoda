package batagodax

import (
	"regexp"
	"testing"
)

// TestCompile tests the regex compilation
func TestCompile(t *testing.T) {
	for _, rule := range Rules {
		_, err := regexp.Compile(rule.Trigger)
		if err != nil {
			t.Errorf("Cannot compile %s", rule.Trigger)
		}
	}
}
