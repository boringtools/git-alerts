package secrets

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/boringtools/git-alerts/logger"
	"github.com/go-git/go-git/v5"
)

func CreateDirectory(dirPath string) {
	cmd := exec.Command("mkdir", dirPath)
	_, err := cmd.Output()

	if err != nil {
		logger.LogERRP("Error in running the command : ", err.Error())
	}
}

func RemoveDirectory(dirPath string) {
	cmd := exec.Command("rm", "-rf", dirPath)
	_, err := cmd.Output()

	if err != nil {
		logger.LogERRP("Error in running the command : ", err.Error())
	}
}

func CloneRepo(url, dir string) {
	_, errClone := git.PlainClone(dir, false, &git.CloneOptions{
		URL:      url,
		Progress: nil,
	})

	if errClone != nil {
		logger.LogERRP("Error in cloning repository : ", errClone)
	}
}

func RunTruffleHog(repoURL string) {
	var tf *exec.Cmd

	_, checkTf := exec.LookPath("trufflehog")

	if checkTf != nil {
		logger.LogERR("Trufflehog is not installed in your machine")
		os.Exit(1)
	}

	if os.Getenv("trufflehog") == "true" {
		tf = exec.Command("trufflehog", "git", repoURL)
	} else {
		tf = exec.Command("trufflehog", "git", repoURL, "--only-verified")
	}

	op, errTf := tf.Output()

	if errTf != nil {
		logger.LogERRP("Error in running the command", errTf.Error())
	} else {
		if string(op) != "" {
			fmt.Println(string(op))
		}
	}
}

func RunGitleaks(repoPath string) {
	_, checkTf := exec.LookPath("gitleaks")

	if checkTf != nil {
		logger.LogERR("Gitleaks is not installed in your machine")
		os.Exit(1)
	}

	gl := exec.Command("gitleaks", "git", repoPath, "-v")
	op, errGl := gl.CombinedOutput()

	if exitError, ok := errGl.(*exec.ExitError); ok && exitError.ExitCode() == 1 {
		fmt.Println(string(op))
		return
	}

	if errGl != nil {
		logger.LogERRP("Error in running the command : ", errGl)
	}
}
