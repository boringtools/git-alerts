package common

import (
	"os"
	"os/exec"

	"github.com/boringtools/git-alerts/internal/ui"
)

func ValidateFlags() {
	if SlackNotification {
		_, isSlackHook := os.LookupEnv("SLACK_HOOK")

		if !isSlackHook {
			ui.PrintError("SLACK_HOOK is not configured in ENV variable")
			os.Exit(1)
		}
	}

	if TrufflehogScan || TrufflehogVerifiedScan {
		_, checkTf := exec.LookPath("trufflehog")

		if checkTf != nil {
			ui.PrintError("trufflehog is not installed in your machine")
		}
	}

	if GitleaksScan {
		_, checkTf := exec.LookPath("gitleaks")

		if checkTf != nil {
			ui.PrintError("gitleaks is not installed in your machine")
		}

		CloneDirectoryPath = ReportPath + CloneDirectoryName

		_, errDirExists := os.Stat(CloneDirectoryPath)

		if os.IsNotExist(errDirExists) {
			CreateDirectory(CloneDirectoryPath)
		} else {
			RemoveDirectory(CloneDirectoryPath)
			CreateDirectory(CloneDirectoryPath)
		}
	}
}
