package clush

import (
	"bytes"
	"os/exec"
)

// CommandReturn is a ran command with its stdout and stderr (once finished)
type CommandReturn struct {
	Stdout string
	Stderr string
}

// execute runs the clush command with given args
func execute(args ...string) (*CommandReturn, error) {
	// Get binary path
	path, err := exec.LookPath("clush")
	if err != nil {
		return nil, err
	}

	// Generate command buffers for stdout/stderr
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// Generate command and execute it
	cmd := exec.Command(path, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()

	// Generate command return
	ret := &CommandReturn{
		Stdout: stdout.String(),
		Stderr: stderr.String(),
	}

	return ret, err
}

// RunOnGroup runs the given command on given clush group
func RunOnGroup(group, command string) (*CommandReturn, error) {
	return execute("-g", group, command)
}

// Version returns the installed clush version
func Version() (string, error) {
	ret, err := execute("--version")
	if err != nil {
		return "", err
	}

	return ret.Stdout, nil
}
