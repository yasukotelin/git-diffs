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
	"github.com/yasukotelin/scrlib"
)

func main() {
	app := cli.NewApp()
	app.Name = "git-diffs"
	app.Version = "1.5.0"
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
		printFiles(stagedFiles, unstagedFiles)
		totalLen := len(stagedFiles) + len(unstagedFiles)
		if totalLen == 0 {
			return nil
		}

		selNum, err := askToSelectFile(stagedFiles, unstagedFiles)
		if err != nil {
			return err
		}
		if selNum == 0 {
			break
		}
		fmt.Println()
		err = runDiff(selNum, stagedFiles, unstagedFiles)
		if err != nil {
			return err
		}

		fmt.Println()
		fmt.Print("(Enter)")
		fmt.Scanln()

		scrlib.Clear()
	}
	return nil
}

func printFiles(staged, unstaged []gitlib.DiffFile) {
	// print steged files
	fmt.Println("Staged files:")
	fmt.Println()
	stagedFilesLen := len(staged)
	if stagedFilesLen == 0 {
		fmt.Println("\tno staged files.")
	}
	for i, file := range staged {
		color.Green.Printf("\t[%d]\t%v\t%v\n", i+1, file.Status, file.Path)
	}
	fmt.Println()

	// print unstaged files
	fmt.Println("Unstaged files:")
	fmt.Println()
	unstagedFilesLen := len(unstaged)
	if unstagedFilesLen == 0 {
		fmt.Println("\tno unstaged files.")
	}
	for i, file := range unstaged {
		color.Red.Printf("\t[%d]\t%v\t%v\n", i+1+stagedFilesLen, file.Status, file.Path)
	}
	fmt.Println()
}

func askToSelectFile(stagedFiles []gitlib.DiffFile, unstagedFiles []gitlib.DiffFile) (int, error) {
	fmt.Print("Select number (empty is cancel) => ")

	var selNumStr string
	fmt.Scanln(&selNumStr)

	if selNumStr == "" {
		return 0, nil
	}
	selNum, err := strconv.Atoi(selNumStr)
	if err != nil {
		return 0, errors.New("your input is not number")
	}
	if selNum > len(stagedFiles)+len(unstagedFiles) || selNum < 1 {
		return 0, errors.New("your input is out of range numbers")
	}

	return selNum, nil
}

func runDiff(number int, stagedFiles, unstagedFiles []gitlib.DiffFile) error {
	var err error
	if number <= len(stagedFiles) {
		// User selected staged file number
		err = gitlib.RunDiff(stagedFiles[number-1].Path, true)
	} else {
		// User selected unstaged file number
		err = gitlib.RunDiff(unstagedFiles[number-len(stagedFiles)-1].Path, false)
	}
	if err != nil {
		return err
	}

	return nil
}

func askToContinue() bool {
	fmt.Print("Continue[Enter] or Quit(q): ")
	var input string
	fmt.Scanln(&input)
	if input == "q" || input == "quit" {
		return false
	}
	return true
}
