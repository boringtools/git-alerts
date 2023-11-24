# GitAlerts

## What problem does it solve?

GitHub repositories created under any organization can be controlled by the GitHub administrators. However any repository created under an organization's user account is not controllable unless the organisation has adopted the GitHub enterprise-managed user (EMU) model.

Any public repository under the organization's user account that was created accidentally or for testing purposes could leak secrets, internal information, code etc. GitAlerts helps you detect and monitor such cases

### Example

> Can be controlled by the administrator `https://github.com/<org>/<org-repo-name>`

> Can't be controlled by the administrator `https://github.com/<org-user>/<org-user-repo-name>`

## Setup

Setup GitHub personal access token (PAT) as the environment variable

```commandline
export GITHUB_PAT=YOUR_GITHUB_PAT
```

## Dependencies

```go
go mod tidy
```

## Usage

Scan GitHub repositories belonging to your organization users

```go
go run . scan --org your-org-name
```

Monitor new public repositories being created by your organization users

```go
go run . monitor --org your-org-name
```

Monitor new public repositories being created by your organization users with slack notification

```go
go run . monitor --org your-org-name --slack-alert
```

Setup slack webhook token as the environment variable

```commandline
export SLACK_HOOK=SLACK_WEBHOOK_URL
```

Scan and generate report with custom path

```go
go run . scan --org your-org-name --report-path /your/file/path/
```

## Documentation

[docs](https://github.com/c0d3G33k/git-alert/tree/main/docs)
> Please feel to reach out for any feedback and suggestions
