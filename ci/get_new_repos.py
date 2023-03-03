from config.conf import *
from common.common import *
from common.json_diff import *

def get_new_repos():
    if MONITOR_MODE:

        new_added_org_repos = get_json_diff(ORG_REPO_FILE_PATH,ORG_REPO_FILE_PATH_DEFAULT,"repository")
        new_added_org_users_repos = get_json_diff(ORG_USER_REPO_FILE_PATH,ORG_USER_REPO_FILE_PATH_DEFAULT,"repository")

        new_repositories = []

        if not new_added_org_repos is None:
            new_repositories = new_repositories + new_added_org_repos

        if not new_added_org_users_repos is None:
            new_repositories = new_repositories + new_added_org_users_repos

        if not new_repositories is None:
           return new_repositories