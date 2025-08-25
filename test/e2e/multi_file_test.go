package e2e

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strings"
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
	wants := make(map[string]struct{})

	for _, tn := range testNodes {
		file, err := createFile(tn.input)
		if err != nil {
			t.Fatal("failed to create test file:", err)
		}
		defer os.Remove(file.Name())
		testFilenames = append(testFilenames, file.Name())
		wants[fmt.Sprintf("%s %s", tn.output, file.Name())] = struct{}{}
	}

	wants[fmt.Sprintf("%s total", testTotals)] = struct{}{}

	cmd := getCommand(t, testFilenames...)

	stdout := &bytes.Buffer{}
	cmd.Stdout = stdout

	if err := cmd.Run(); err != nil {
		t.Fatal("failed to run command:", err)
	}

	scanner := bufio.NewScanner(stdout)

	for scanner.Scan() {
		line := scanner.Text()
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		if _, ok := wants[line]; ok {
			delete(wants, line)
		} else {
			t.Log("unexpected line:", line)
			t.Fail()
		}
	}

	if len(wants) != 0 {
		t.Fatal("desired output not matched:", wants)
	}

}
