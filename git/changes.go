package git

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func gitRoot() string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	dat, err := cmd.Output()
	if err != nil {
		Die("Could not find git root: %s", err)
	}

	return strings.TrimSpace(string(dat))
}

func ChangedFiles(commitRange string) []string {
	root := gitRoot()

	cmd := exec.Command("git", "diff-tree", "--no-commit-id", "--name-only", "-r", commitRange)
	fmt.Println(cmd)
	dat, err := cmd.Output()
	if err != nil {
		Die("Could not run git diff-tree: %v", err)
	}
	files := strings.Split(string(dat), "\n")
	var res []string
	for _, f := range files {
		f = strings.TrimSpace(f)
		if len(f) == 0 {
			continue
		}
		res = append(res, filepath.Join(root, f))
	}
	return res
}

func Die(s string, i ...interface{}) { panic(fmt.Sprintf(s, i...)) }
