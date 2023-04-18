package git

import (
	"os/exec"
)

type CommandRunner func(args ...string) ([]byte, error)

func RunGit(args ...string) (content []byte, err error) {
	// log.Debug().Strs("args", args).Msg("Running git command")
	cmd := exec.Command("git", args...)
	out, _ := cmd.Output()

	return out, nil
}
