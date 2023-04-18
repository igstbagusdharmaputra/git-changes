package main

import (
	"fmt"

	"github.com/igstbagusdharmaputra/git-changes/git"
)

func main() {
	files, _ := git.ChangedFiles(git.RunGit, "HEAD ^origin/main")
	fmt.Printf("main %s", files)
}
