from config.conf import *

GITHUB_PAT = validate_env_variable("GITHUB_PAT")
GITHUB_REQUEST_HEADERS = {"Authorization": f"token {GITHUB_PAT}"}

users_directory = {"usernames": []}
organization_directory = {"repository": [], "fork": [], "created": [], "updated": [], "pushed": []}
organization_users_directory = {"repository": [], "fork": [], "created": [], "updated": [], "pushed": []}
github_request_params = {"page": "1", "per_page": "100"}

github_url_org = f"https://api.github.com/orgs/{ORGANIZATION}"
github_url_user = f"https://api.github.com/users"

