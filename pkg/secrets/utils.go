package secrets

import (
	"fmt"
	"os/exec"

	"github.com/boringtools/git-alerts/pkg/common"
	"github.com/go-git/go-git/v5"
)

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

	if common.TrufflehogVerifiedScan {
		tf = exec.Command("trufflehog", "git", repoURL, "--only-verified")
	} else {
		tf = exec.Command("trufflehog", "git", repoURL)
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
	gitleaksScan := exec.Command("gitleaks", "git", repoPath, "-v")
	output, errGitleaksScan := gitleaksScan.CombinedOutput()

	if exitError, ok := errGitleaksScan.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
		fmt.Println(string(output))
	}

	if errGitleaksScan != nil {
		return errGitleaksScan
	}

	return nil
}
