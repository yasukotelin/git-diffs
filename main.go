package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/gookit/color"
	"github.com/urfave/cli"
	"github.com/yasukotelin/gitlib"
)

func main() {
	app := cli.NewApp()
	app.Name = "git-diffs"
	app.Version = "1.4.2"
	app.Description = "The git subcommand that is diff files selector."
	app.Action = mainAction

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func mainAction(c *cli.Context) error {
	// get staged files.
	stagedFiles, err := gitlib.GetDiffFiles(true)
	if err != nil {
		return err
	}
	// get unstaged files.
	unstagedFiles, err := gitlib.GetDiffFiles(false)
	if err != nil {
		return err
	}

	for {
		isContinue, err := askToSelectFile(stagedFiles, unstagedFiles)
		if err != nil {
			return err
		}
		if !isContinue {
			break
		}
		fmt.Println()
	}

	return nil
}

func askToSelectFile(stagedFiles []gitlib.DiffFile, unstagedFiles []gitlib.DiffFile) (isContinue bool, err error) {
	fmt.Println("Staged files:")
	fmt.Println()
	stagedFilesLen := len(stagedFiles)
	if stagedFilesLen == 0 {
		fmt.Println("\tno staged files.")
	}
	for i, file := range stagedFiles {
		color.Green.Printf("\t[%d]\t%v\t%v\n", i+1, file.Status, file.Path)
	}
	fmt.Println()

	fmt.Println("Unstaged files:")
	fmt.Println()
	unstagedFilesLen := len(unstagedFiles)
	if unstagedFilesLen == 0 {
		fmt.Println("\tno unstaged files.")
	}
	for i, file := range unstagedFiles {
		color.Red.Printf("\t[%d]\t%v\t%v\n", i+1+stagedFilesLen, file.Status, file.Path)
	}
	fmt.Println()

	if stagedFilesLen+unstagedFilesLen == 0 {
		return false, nil
	}

	fmt.Print("Select number (empty is cancel) => ")

	var selNumStr string
	fmt.Scanln(&selNumStr)

	if selNumStr == "" {
		return false, nil
	}
	selNum, err := strconv.Atoi(selNumStr)
	if err != nil {
		return false, errors.New("your input is not number")
	}
	if selNum > stagedFilesLen+unstagedFilesLen || selNum < 1 {
		return false, errors.New("your input is out of range numbers")
	}

	fmt.Println()
	if selNum <= stagedFilesLen {
		// User selected staged file number
		err = gitlib.RunDiff(stagedFiles[selNum-1].Path, true)
	} else {
		// User selected unstaged file number
		err = gitlib.RunDiff(unstagedFiles[selNum-stagedFilesLen-1].Path, false)
	}
	if err != nil {
		return false, err
	}

	return true, nil
}
