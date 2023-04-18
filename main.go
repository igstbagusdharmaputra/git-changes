package main

import (
	"fmt"

	"github.com/igstbagusdharmaputra/git-changes/git"
)

func main() {
	files := git.ChangedFiles("main")
	fmt.Println(files)
}
