package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/urfave/cli"
	"github.com/yasukotelin/git-diffs/git"
	"github.com/gookit/color"
)

func main() {
	app := cli.NewApp()
	app.Name = "git-diffs"
	app.Version = "1.3.0"
	app.Description = "The git subcommand that is diff files selector."
	app.Action = mainAction

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func mainAction(c *cli.Context) error {
	// get staged files.
	stagedFiles, err := git.DiffFiles(true)
	if err != nil {
		return err
	}
	// get unstaged files.
	unstagedFiles, err := git.DiffFiles(false)
	if err != nil {
		return err
	}

	fmt.Println("Staged files:")
	fmt.Println()
	stagedFilesLen := len(stagedFiles)
	if stagedFilesLen == 0 {
		fmt.Println("no staged files.")
	}
	for i, file := range stagedFiles {
		color.Green.Printf("\t[%d]\t%v\t%v\n", i+1, file.Status, file.Path)
	}
	fmt.Println()

	fmt.Println("Unstaged files:")
	fmt.Println()
	unstagedFilesLen := len(unstagedFiles)
	if unstagedFilesLen == 0 {
		fmt.Println("no unstaged files.")
	}
	for i, file := range unstagedFiles {
		color.Red.Printf("\t[%d]\t%v\t%v\n", i+1+stagedFilesLen, file.Status, file.Path)
	}
	fmt.Println()

	if (stagedFilesLen + unstagedFilesLen == 0) {
		return nil
	}
	
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
	if selNum > stagedFilesLen+unstagedFilesLen || selNum < 1 {
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
