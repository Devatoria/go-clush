package clush

import (
	"bytes"
	"os/exec"
	"strings"
)

// execute runs the clush command with given args
func execute(args ...string) (string, string, error) {
	// Get binary path
	path, err := exec.LookPath("clush")
	if err != nil {
		return "", "", err
	}

	// Generate command buffers for stdout/stderr
	var stdout bytes.Buffer
	var stderr bytes.Buffer

	// Generate command and execute it
	cmd := exec.Command(path, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()

	return stdout.String(), stderr.String(), err
}

// RunOnGroup runs the given command on given clush group
func RunOnGroup(group, command string) (string, string, error) {
	return execute("-g", group, command)
}

// RunOnNodes runs the given command on given nodes, and excludes some nodes if provided
func RunOnNodes(nodes []string, excluded []string, command string) (string, string, error) {
	// Add -w option with given nodes
	// and -x option with excluded nodes if provided
	exec := []string{"-w", strings.Join(nodes, ",")}
	if len(excluded) > 0 {
		exec = append(exec, "-x", strings.Join(excluded, ","))
	}

	// Append command
	exec = append(exec, command)

	return execute(exec...)
}

// Version returns the installed clush version
func Version() (string, error) {
	stdout, _, err := execute("--version")
	if err != nil {
		return "", err
	}

	return stdout, nil
}
