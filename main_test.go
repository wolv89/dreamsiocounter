package main_test

import (
	"testing"

	counter "github.com/wolv89/dreamsiocounter"
)

func TestCountWords(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 words",
			input: "one two three four five",
			wants: 5,
		},
		{
			name:  "empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "single space",
			input: " ",
			wants: 0,
		},
		{
			name:  "new line",
			input: "one two three\nfour five",
			wants: 5,
		},
		{
			name:  "multiple spaces",
			input: "What if  maybe I hit spacebar    too many    times??",
			wants: 9,
		},
		{
			name:  "prefix spaces",
			input: "   hello",
			wants: 1,
		},
		{
			name:  "suffix spaces",
			input: "world    ",
			wants: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result := counter.CountWords([]byte(tc.input))

			if result != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, result)
				t.Fail()
			}

		})
	}

}
