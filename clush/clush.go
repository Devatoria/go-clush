package clush

import (
	"bytes"
	"os/exec"
	"strconv"
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
	return RunOnGroupWithFanout(group, 0, command)
}

// RunOnGroupWithFanout runs the given command on the given clush group, applying
// the given fanout
func RunOnGroupWithFanout(group string, fanout int, command string) (string, string, error) {
	args := []string{"-g", group}
	if fanout > 0 {
		args = append(args, "-f", strconv.Itoa(fanout))
	}

	args = append(args, command)

	return execute(args...)
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
