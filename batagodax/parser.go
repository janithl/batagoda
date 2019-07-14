package batagodax

import (
	"math/rand"
	"strings"
)

// Respond responds to a prompt
func Respond(prompt string) string {
	prompt = cleanString(prompt)

	// Iterates through Rules, tries to match to prompt
	for _, rule := range Rules {
		for _, trigger := range rule.Triggers {
			if strings.Contains(prompt, cleanString(trigger)) {
				return randChoice(rule.Responses)
			}
		}
	}

	// If all else fails, return a fallback
	return randChoice(Fallbacks)
}

// cleanString cleans out special unicode chars from strings
func cleanString(str string) string {
	const (
		ZWSP   = '\u200B'
		ZWNBSP = '\uFEFF'
		ZWJ    = '\u200D'
		ZWNJ   = '\u200C'
	)

	r := strings.NewReplacer(
		string(ZWSP), "",
		string(ZWNBSP), "",
		string(ZWJ), "",
		string(ZWNJ), "",
	)

	return r.Replace(str)
}

// randChoice returns a random element in an (string) array.
func randChoice(list []string) string {
	randIndex := rand.Intn(len(list))
	return list[randIndex]
}
