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
	for i, file := range unstagedFiles {
		fmt.Printf("[%d] %v\n", i+1, file.Text)
	}

	// files, err := execDiffNameOnly()
	// if err != nil {
	// 	return err
	// }

	// for i, f := range files {
	// 	fmt.Printf("[%d] %s\n", i+1, f)
	// }

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
	fmt.Println(selNum)
	// if selNum > len(files) || selNum < 1 {
	// 	return errors.New("your input is out of range numbers")
	// }
	// execDiff(files[selNum-1])

	return nil
}
