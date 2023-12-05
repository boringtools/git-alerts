package secrets

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/boringtools/git-alerts/logger"
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
	cl := exec.Command("git", "clone", url, dir)
	_, errCl := cl.Output()

	if errCl != nil {
		logger.LogERRP("Error in running the command", errCl.Error())
	}
}

func RunTruffleHog(file string) {
	_, checkTf := exec.LookPath("trufflehog")

	if checkTf != nil {
		logger.LogERR("Trufflehog is not installed in your machine")
		os.Exit(1)
	}

	tf := exec.Command("trufflehog", "git", file, "--only-verified")
	op, errTf := tf.Output()

	if errTf != nil {
		logger.LogERRP("Error in running the command", errTf.Error())
	} else {
		if string(op) != "" {
			fmt.Println(string(op))
		}
	}
}
