package e2e

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func getCommand(t *testing.T, args ...string) *exec.Cmd {

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal("could not get working directory:", err)
	}

	path := filepath.Join(dir, binName)

	return exec.Command(path, args...)

}

func createFile(content string) (*os.File, error) {

	file, err := os.CreateTemp("", "counter-test-*")
	if err != nil {
		return nil, fmt.Errorf("unable to create temporary file for testing: %w", err)
	}

	_, err = file.WriteString(content)
	if err != nil {
		return nil, fmt.Errorf("unable to write to temporary file for testing: %w", err)
	}

	err = file.Close()
	if err != nil {
		return nil, fmt.Errorf("unable to close temporary file for testing: %w", err)
	}

	return file, nil

}

func TestStdin(t *testing.T) {

	cmd := getCommand(t)
	output := &bytes.Buffer{}

	cmd.Stdin = strings.NewReader("one two three\n")
	cmd.Stdout = output

	if err := cmd.Run(); err != nil {
		t.Fatal("failed to run command:", err)
	}

	wants := " 1 3 14\n\n"

	if wants != output.String() {
		t.Log("stdout is not correct wants: '", wants, "' got: '", output.String(), "'")
		t.Fail()
	}

}

func TestSingleFile(t *testing.T) {

	file, err := os.CreateTemp("", "counter-test-*")
	if err != nil {
		t.Fatal("unable to create temporary file for testing:", err)
	}

	defer os.Remove(file.Name())

	_, err = file.WriteString("foo bar baz\nbaz bar foo\none two three\n")
	if err != nil {
		t.Fatal("unable to write to temporary file:", err)
	}

	defer file.Close()

	cmd := getCommand(t, file.Name())
	output := &bytes.Buffer{}

	cmd.Stdout = output

	if err := cmd.Run(); err != nil {
		t.Fatal("failed to run command:", err)
	}

	wants := fmt.Sprintf(" 3 9 38 %s\n\n", file.Name())

	if wants != output.String() {
		t.Log("stdout is not correct wants: '", wants, "' got: '", output.String(), "'")
		t.Fail()
	}

}

func TestNoExists(t *testing.T) {

	cmd := getCommand(t, "missing-file.txt")

	stderr := &bytes.Buffer{}
	stdout := &bytes.Buffer{}

	cmd.Stderr = stderr
	cmd.Stdout = stdout

	wantsStderr := "counter: open missing-file.txt: no such file or directory\n"
	wantsStdout := "\n"

	err := cmd.Run()

	if err == nil {
		t.Log("command ran successfully, expecting failure")
		t.Fail()
	}

	if err.Error() != "exit status 1" {
		t.Log("unexpected error:", err)
		t.Fail()
	}

	if stderr.String() != wantsStderr {
		t.Log("stderr is not correct wants: '", wantsStderr, "' got: '", stderr.String(), "'")
		t.Fail()
	}

	if stdout.String() != wantsStdout {
		t.Log("stdout is not correct wants: '", wantsStdout, "' got: '", stdout.String(), "'")
		t.Fail()
	}

}
