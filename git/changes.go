package git

import (
	"fmt"
	"path/filepath"
	"strings"
)

func gitRoot(cmd CommandRunner) (string, error) {

	out, err := cmd("rev-parse", "--show-toplevel")
	if err != nil {
		return "", fmt.Errorf("failed to  show toplevel from git: %w", err)
	}
	dat := out

	return strings.TrimSpace(string(dat)), nil
}

func ChangedFiles(cmd CommandRunner, commitRange string) ([]string, error) {
	root, _ := gitRoot(cmd)

	out, err := cmd("diff-tree", "--no-commit-id", "--name-only", "-r", commitRange)
	if err != nil {
		return nil, fmt.Errorf("failed to get the list of modified files from git: %w", err)
	}

	dat := out

	files := strings.Split(string(dat), "\n")

	var res []string
	for _, f := range files {
		f = strings.TrimSpace(f)
		if len(f) == 0 {
			continue
		}
		res = append(res, filepath.Join(root, f))
	}
	return res, nil
}
