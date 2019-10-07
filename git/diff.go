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
	Status   string
	Path     string
	IsStaged bool
}

func DiffFiles(isStaged bool) ([]DiffFile, error) {
	nameOnly, err := DiffNameOnly(isStaged)
	if err != nil {
		return nil, err
	}
	nameStatus, err := DiffNameStatus(isStaged)
	if err != nil {
		return nil, err
	}

	// len(nameStatuses) equal len(nameOnlies)
	rowLen := len(nameOnly)

	diffFile := make([]DiffFile, rowLen)
	for i := 0; i < rowLen; i++ {
		path := nameOnly[i]
		status := strings.Fields(nameStatus[i])

		diffFile[i] = DiffFile{
			Status: status[0],
			Path: path,
			IsStaged: isStaged,
		}
	}

	return diffFile, nil
}

// DiffNameStatus runs `git diff --name-status` git command.
// If isStaged is true, add `--staged`.
func DiffNameStatus(isStaged bool) ([]string, error) {
	var cmd *exec.Cmd
	if isStaged {
		cmd = exec.Command("git", "diff", "--staged", "--name-status")
	} else {
		exec.Command("git", "add", "-A", "-N").Run()
		cmd = exec.Command("git", "diff", "--name-status")
	}
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	rows := strings.Split(string(out), "\n")
	// Remove the latest empty row.
	return rows[0 : len(rows)-1], nil
}

// DiffNameOnly runs `git diff --name-only` git commnad.
// If isStaged is true, add `--staged`.
func DiffNameOnly(isStaged bool) ([]string, error) {
	var cmd *exec.Cmd
	if isStaged {
		cmd = exec.Command("git", "diff", "--staged", "--name-only")
	} else {
		exec.Command("git", "add", "-A", "-N").Run()
		cmd = exec.Command("git", "diff", "--name-only")
	}
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	rows := strings.Split(string(out), "\n")
	// Remove the latest empty row.
	return rows[0 : len(rows)-1], nil
}
