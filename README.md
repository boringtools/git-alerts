# GitAlert

## What problem does it solve?

GitHub repositories created under any organization can be controlled by the administrators. But any repository created under an organization user account is not controllable unless the organisation has adopted the GitHub enterprise-managed user model. 

Any public repository created under the organization user account that was created for any testing could leak secrets, internal information, code etc.

### Example

- https://github.com/<org>/<org-repo-name>
- https://github.com/<org-user>/<org-user-repo-name>

`git-alert` helps you to detect and monitor public repositories creation under the organization and organization users as well.

## Setup

Setup GitHub personal access token as the environment variable

```commandline
export GITHUB_PAT=YOUR_GITHUB_PAT
```
## Dependencies

```commandline
pip3 install -r requirements.txt
```
## Usage

> Find all public GitHub repositories belonging to an organization and organization users

```commandline
python3 main.py -o your-organization-name
```

> Run in monitor mode to detect new public repositories

```commandline
python3 main.py -o your-organization-name -m True
```

> Run in monitor mode to detect new public repositories with slack notifications

```commandline
python3 main.py -o your-organization-name -m True -s True
```
Setup slack webhook token as the environment variable
```commandline
export SLACK_WEBHOOK=SLACK_WEBHOOK_TOKEN
```
[docs](https://github.com/boringtools/git-alerts/tree/main/docs)
> For future work & support, please check the issues created

> Please feel to reach out for any feedback and suggestions
