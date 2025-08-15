package main_test

import (
	"strings"
	"testing"

	counter "github.com/wolv89/dreamsiocounter"
)

func TestGetCounts(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		wants counter.Counts
	}{
		{
			name:  "5 words",
			input: "one two three four five\n",
			wants: counter.Counts{
				Lines: 1,
				Words: 5,
				Bytes: 24,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			result := counter.GetCounts(r)

			if result != tc.wants {
				t.Logf("expected: %v got: %v", tc.wants, result)
				t.Fail()
			}

		})
	}

}

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
			r := strings.NewReader(tc.input)
			result := counter.CountWords(r)

			if result != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, result)
				t.Fail()
			}

		})
	}

}

func TestCountLines(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 words",
			input: "one two three four five\n",
			wants: 1,
		},
		{
			name:  "empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "single newline",
			input: "\n",
			wants: 1,
		},
		{
			name:  "two lines",
			input: "one two three\nfour five\n",
			wants: 2,
		},
		{
			name:  "multiple lines",
			input: "What if\n\nmaybe I hit enter\n\n\ntoo many\ntimes??",
			wants: 6,
		},
		{
			name:  "prefix spaces",
			input: "\n\nhello",
			wants: 2,
		},
		{
			name:  "suffix spaces",
			input: "world\n\n\n",
			wants: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			result := counter.CountLines(r)

			if result != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, result)
				t.Fail()
			}

		})
	}

}

func TestCountBytes(t *testing.T) {

	testCases := []struct {
		name  string
		input string
		wants int
	}{
		{
			name:  "5 words",
			input: "one two three four five",
			wants: 23,
		},
		{
			name:  "empty input",
			input: "",
			wants: 0,
		},
		{
			name:  "all spaces",
			input: "         ",
			wants: 9,
		},
		{
			name:  "newlines and tabs",
			input: "one\ntwo\tthree",
			wants: 13,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := strings.NewReader(tc.input)
			result := counter.CountBytes(r)

			if result != tc.wants {
				t.Logf("expected: %d got: %d", tc.wants, result)
				t.Fail()
			}

		})
	}

}

var benchData = []string{
	"this is a test data string\nthat runs across\nmultiple lines.\n",
	"one two three\nfour\nfive\nsix\nseven\neight",
	"this is a weird\n\n\n\n\n\n\n       string.",
}

func BenchmarkGetCounts(b *testing.B) {
	for i := range b.N {
		data := benchData[i%len(benchData)]
		r := strings.NewReader(data)
		counter.GetCounts(r)
	}
}

func BenchmarkGetCountsSinglePass(b *testing.B) {
	for i := range b.N {
		data := benchData[i%len(benchData)]
		r := strings.NewReader(data)
		counter.GetCountsSinglePass(r)
	}
}
