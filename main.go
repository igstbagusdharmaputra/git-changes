package main

import (
	"flag"
	"fmt"

	"github.com/igstbagusdharmaputra/git-changes/git"
)

func main() {
	commitRange := flag.String("CR", "main", "value commit..commit or <branch name>")
	flag.Parse()

	files := git.ChangedFiles(*commitRange)
	fmt.Println(files)
}
