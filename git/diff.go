package git

import (
	"os"
	"os/exec"
	"strings"
)

func Diff(path string, isStaged bool) error {
	var cmd *exec.Cmd
	if isStaged {
		cmd = exec.Command("git", "diff", "--staged", path)
	} else {
		cmd = exec.Command("git", "diff", path)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

type DiffFile struct {
	Text     string
	Status   string
	Path     string
	IsStaged bool
}

func DiffNameStatus(isStaged bool) ([]DiffFile, error) {
	var cmd *exec.Cmd
	if isStaged {
		cmd = exec.Command("git", "diff", "--staged", "--name-status")
	} else {
		cmd = exec.Command("git", "diff", "--name-status")
	}
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	rows := strings.Split(string(out), "\n")
	// Remove the latest empty row.
	rows = rows[0 : len(rows)-1]

	// Divide the status(M, D etc) and file path
	diffFile := make([]DiffFile, len(rows))
	for i, row := range rows {
		splited := strings.Fields(row)
		diffFile[i] = DiffFile{
			Text:     row,
			Status:   splited[0],
			Path:     splited[1],
			IsStaged: isStaged,
		}
	}

	return diffFile, nil
}
