package batagodax

import (
	"math/rand"
	"strings"
)

// Respond responds to a prompt
func Respond(prompt string) string {
	// Iterates through Rules, tries to match to prompt
	for _, rule := range Rules {
		for _, trigger := range rule.Triggers {
			if strings.Contains(prompt, trigger) {
				return randChoice(rule.Responses)
			}
		}
	}

	// If all else fails, return a fallback
	return randChoice(Fallbacks)
}

// randChoice returns a random element in an (string) array.
func randChoice(list []string) string {
	randIndex := rand.Intn(len(list))
	return list[randIndex]
}
