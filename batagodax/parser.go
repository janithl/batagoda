package batagodax

import (
	"math/rand"
	"regexp"
)

// Respond responds to a prompt
func Respond(prompt string) string {
	// Iterates through Rules, tries to match to prompt
	for _, rule := range Rules {
		re, err := regexp.Compile(rule.Trigger)
		if err != nil {
			// If there's an error, return a fallback
			return randChoice(Fallbacks)
		}

		// Try to find matches for the trigger pattern
		matches := re.FindStringSubmatch(prompt)
		if len(matches) > 0 {
			return randChoice(rule.Responses)
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
