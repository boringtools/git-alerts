package secrets

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/go-git/go-git/v5"
)

func CreateDirectory(dirPath string) error {
	cmd := exec.Command("mkdir", dirPath)
	_, err := cmd.Output()

	if err != nil {
		return err
	}
	return nil
}

func RemoveDirectory(dirPath string) error {
	cmd := exec.Command("rm", "-rf", dirPath)
	_, err := cmd.Output()

	if err != nil {
		return err
	}

	return nil
}

func CloneRepo(url, dir string) error {
	_, errClone := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      url,
		Progress: nil,
	})

	if errClone != nil {
		return errClone
	}
	return nil
}

func RunTruffleHog(repoURL string) error {
	var tf *exec.Cmd

	_, checkTf := exec.LookPath("trufflehog")

	if checkTf != nil {
		return fmt.Errorf("trufflehog is not installed in your machine")
	}

	if os.Getenv("trufflehog") == "true" {
		tf = exec.Command("trufflehog", "git", repoURL)
	} else {
		tf = exec.Command("trufflehog", "git", repoURL, "--only-verified")
	}

	op, errTf := tf.Output()

	if errTf != nil {
		return errTf
	} else {
		if string(op) != "" {
			fmt.Println(string(op))
		}
	}

	return nil
}

func RunGitleaks(repoPath string) error {
	_, checkTf := exec.LookPath("gitleaks")

	if checkTf != nil {
		return fmt.Errorf("gitleaks is not installed in your machine")
	}

	gl := exec.Command("gitleaks", "git", repoPath, "-v")
	op, errGl := gl.CombinedOutput()

	if exitError, ok := errGl.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
		fmt.Println(string(op))
	}

	if errGl != nil {
		return errGl
	}

	return nil
}
