package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"errors"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "git-diffs"
	app.Version = "0.1.0"
	app.Description = "The git subcommand that is diff files selector."
	app.Action = mainAction

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func mainAction(c *cli.Context) error {
	files, err := execDiffNameOnly()
	if err != nil {
		return err
	}

	for i, f := range files {
		fmt.Printf("[%d] %s\n", i+1, f)
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
	if selNum > len(files) || selNum < 1 {
		return errors.New("your input is out of range numbers")
	}
	execDiff(files[selNum-1])

	return nil
}
