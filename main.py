from github.get_org_users import *
from github.get_org_repositories import *
from github.get_org_users_repositories import *
from common.save_to_csv import *
from common.save_to_json import *
from notification.message import *
from ci.monitor import *

def main():
    org_users = get_org_users()
    save_to_json(org_users, ORG_USERS_FILE_NAME)
    save_to_csv(ORG_USERS_FILE_NAME)

    org_repo = get_org_repositories()
    save_to_json(org_repo, ORG_REPO_FILE_NAME)
    save_to_csv(ORG_REPO_FILE_NAME)

    org_user_repo = get_org_users_repositories()
    save_to_json(org_user_repo, ORG_USERS_REPO_FILE_NAME)
    save_to_csv(ORG_USERS_REPO_FILE_NAME)

    scan_summary()
    monitor()


if __name__ == '__main__':
    main()