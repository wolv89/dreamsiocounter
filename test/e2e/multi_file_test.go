package e2e

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestMultipleFiles(t *testing.T) {

	type testNode struct {
		input, output string
	}

	testNodes := []testNode{
		{
			input:  "one two three four five\n",
			output: " 1 5 24",
		},
		{
			input:  "foo bar baz\n\n",
			output: " 2 3 13",
		},
		{
			input:  "",
			output: " 0 0  0",
		},
	}
	testTotals := " 3 8 37"

	testFilenames := make([]string, 0, len(testNodes))
	var wants string

	for _, tn := range testNodes {
		file, err := createFile(tn.input)
		if err != nil {
			t.Fatal("failed to create test file:", err)
		}
		defer os.Remove(file.Name())
		testFilenames = append(testFilenames, file.Name())
		wants += fmt.Sprintf("%s %s\n", tn.output, file.Name())
	}

	wants += fmt.Sprintf("%s total\n\n", testTotals)

	cmd := getCommand(t, testFilenames...)

	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		t.Fatal("failed to run command:", err)
	}

	res := stdout.String()

	if res != wants {
		t.Fatal("wants:\n", wants, " got:\n", res)
	}

}
