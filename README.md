# GitAlert

GitAlert tool detects and alerts public repositories belonging to an organization and organization users that may leak any secrets along with various misconfigurations

## Setup

Setup GitHub personal access token as environment variable

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
> For future work & support, please check the issues created