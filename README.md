# GitAlerts

[![Go Report Card](https://goreportcard.com/badge/github.com/boringtools/git-alerts)](https://goreportcard.com/report/github.com/boringtools/git-alerts)
![License](https://img.shields.io/github/license/boringtools/git-alerts)
![Release](https://img.shields.io/github/v/release/boringtools/git-alerts)

## What problem does it solve?

GitHub repositories created under any organization can be controlled by the GitHub administrators. However any repository created under an organization's user account is not controllable unless the organisation has adopted the GitHub enterprise-managed user (EMU) model.

Any public repository under the organization's user account that was created accidentally or for testing purposes could leak secrets, internal information, code etc. GitAlerts helps you detect and monitor such cases

### Example

> Can be controlled by the administrator `https://github.com/<org>/<org-repo-name>`

> Can't be controlled by the administrator `https://github.com/<org-user>/<org-user-repo-name>`

## Getting Started

- Download the binary file for your operating system / architecture from the [Official GitHub Releases](https://github.com/boringtools/git-alerts/releases)

- You can also install `git-alerts` using homebrew in MacOS and Linux

```bash
brew tap boringtools/tap
brew install boringtools/tap/git-alerts
```

- Alternatively, build from source

> Ensure $(go env GOPATH)/bin is in your $PATH

```bash
go install github.com/boringtools/git-alerts@main
```

Setup GitHub personal access token [(PAT)](https://github.com/boringtools/git-alerts/blob/main/docs/github.md) as the environment variable

```bash
export GITHUB_PAT=YOUR_GITHUB_PAT
```

## Usage

Scan GitHub repositories belonging to your organization users

```bash
git-alerts scan --org your-org-name
```

Monitor new public repositories being created by your organization users

```bash
git-alerts monitor --org your-org-name
```

Monitor new public repositories being created by your organization users with slack notification

```bash
git-alerts monitor --org your-org-name --slack-alert
```

Setup slack webhook token as the environment variable

```bash
export SLACK_HOOK=SLACK_WEBHOOK_URL
```

Scan and generate report with custom path

```bash
git-alerts scan --org your-org-name --report-path /your/file/path/
```

Scan with secrets detection using Trufflehog
> Ensure trufflehog is installed in your machine

```bash
git-alerts detect --org your-org-name
```

## Documentation

[docs](https://github.com/boringtools/git-alerts/tree/main/docs)
> Please feel to reach out for any feedback and suggestions

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=boringtools/git-alerts&type=Date)](https://star-history.com/#boringtools/git-alerts&Date)
