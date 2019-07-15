package batagodax

import (
	"fmt"
	"testing"
)

// testcase is a struct to hold testcases
type testcase struct {
	Prompt string
	Want   int
}

// TestMatch tests if prompts get the expected bot responses
func TestMatch(t *testing.T) {
	testCases := []testcase{
		{"අපි හුකමු​කො", 0},
		{"තෝ නම් කැරියෙක්", 1},
		{"මම ඔයාට ආදරෙයි", 2},
		{"මම උඹට මනාපයි", 2},
		{"ටත් ලවු සික් ගැහුව කාලයක් තිබ්බ", 2},
		{"ආස කෙල්ලන්ට ද කොල්ලන්ට ද?", 3},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Prompt: %s", tc.Prompt), func(t *testing.T) {
			response := Respond(tc.Prompt)
			got := -1

			for i, rule := range Rules {
				for _, resp := range rule.Responses {
					if response == resp {
						got = i
					}
				}
			}

			if got != tc.Want {
				t.Errorf("got %d; want %d", got, tc.Want)
			}
		})
	}
}
