package main

import (
	"fmt"

	"github.com/igstbagusdharmaputra/git-changes/git"
)

func main() {
	commitRange, _ := git.CommitRange(git.RunGit, "main")
	changes, _ := git.Changes(git.RunGit, commitRange)
	fmt.Println(changes)
}
