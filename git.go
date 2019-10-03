package main

import (
	"os"
	"os/exec"
	"strings"
)

func execDiffNameOnly() ([]string, error) {
	// git add -N すると実際にはaddしていないがaddしたときと同じような結果が得られるようになる
	// つまり、つぎにgit diffしたときに新規作成ファイルもdiff表示されるようになる
	err := exec.Command("git", "add", "-N", "-A").Run()
	if err != nil {
		return nil, err
	}

	out, err := exec.Command("git", "diff", "--name-only").Output()
	if err != nil {
		return nil, err
	}

	files := strings.Split(string(out), "\n")

	// 最後の空行は抜き出して返却
	return files[0:len(files)-1], nil
}

func execDiff(path string) error {
	cmd := exec.Command("git", "diff", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
