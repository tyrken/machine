package main

import (
	"os/exec"
	"testing"
)

func TestVirtualBox(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	runCmd := exec.Command(machineBinary, "create", "-d", "virtualbox", "machine-test-virtualbox")
	out, exitCode, err := runCommandWithOutput(runCmd)
	if err != nil {
		t.Fatal(out, err)
	}
	if exitCode != 0 {
		t.Fatalf("error creating machine: exit code: %d; output: %s", exitCode, out)
	}
}
