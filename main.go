package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli"
	"github.com/yasukotelin/git-diffs/git"
)

func main() {
	app := cli.NewApp()
	app.Name = "git-diffs"
	app.Version = "1.1.0"
	app.Description = "The git subcommand that is diff files selector."
	app.Action = mainAction

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func mainAction(c *cli.Context) error {
	// get staged files.
	stagedFiles, err := git.DiffNameStatus(true)
	if err != nil {
		return err
	}
	fmt.Println("=== Staged files ===")
	stagedFilesLen := len(stagedFiles)
	for i, file := range stagedFiles {
		fmt.Printf("[%d] %v\n", i+1, file.Text)
	}

	fmt.Println()

	// get unstaged files.
	unstagedFiles, err := git.DiffNameStatus(false)
	if err != nil {
		return err
	}
	fmt.Println("=== Unstaged files ===")
	unstagedFilesLen := len(unstagedFiles)
	for i, file := range unstagedFiles {
		fmt.Printf("[%d] %v\n", i+1+stagedFilesLen, file.Text)
	}

	fmt.Println()
	fmt.Print("Select number (empty is cancel) => ")

	var selNumStr string
	fmt.Scanln(&selNumStr)

	if selNumStr == "" {
		return nil
	}
	selNum, err := strconv.Atoi(selNumStr)
	if err != nil {
		return errors.New("your input is not number.")
	}
	if selNum > stagedFilesLen + unstagedFilesLen || selNum < 1 {
		return errors.New("your input is out of range numbers")
	}

	if selNum <= stagedFilesLen {
		// User selected staged file number
		err = git.Diff(stagedFiles[selNum-1].Path, true)
	} else {
		// User selected unstaged file number
		err = git.Diff(unstagedFiles[selNum-stagedFilesLen-1].Path, false)
	}
	if err != nil {
		return err
	}

	return nil
}
