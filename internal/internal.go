package internal

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

// Exec checks for existence of first argument as an executable on the
// system and then runs it with exec.Command.Run  exiting in a way that
// is supported across all architectures that Go supports. The stdin,
// stdout, and stderr are connected directly to that of the calling
// program. Sometimes this is insufficient and the UNIX-specific SysExec
// is preferred. See exec.Command.Run for more about distinguishing
// ExitErrors from others.
func Exec(args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("missing name of executable")
	}
	path, err := exec.LookPath(args[0])
	if err != nil {
		return err
	}
	cmd := exec.Command(path, args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// Has returns true if any of the names are found in the exec.LookPath.
func Has(names ...string) bool {
	for _, name := range names {
		_, err := exec.LookPath(name)
		if err == nil {
			return true
		}
	}
	return false
}

func GOOS() string   { return runtime.GOOS }
func GOARCH() string { return runtime.GOARCH }
