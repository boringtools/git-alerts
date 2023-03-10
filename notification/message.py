from tabulate import tabulate
from common.get_json_length import *
from config.conf import *
from logger.log import *
def scan_summary():
    try:
        org_users_count = get_json_len(ORG_USER_FILE_PATH, 'usernames')
        org_repos_count = get_json_len(ORG_REPO_FILE_PATH, 'repository')
        org_users_repo_count = get_json_len(ORG_USER_REPO_FILE_PATH, 'repository')
        org_upper_name = ORGANIZATION.upper()
        data = [
            [f'TOTAL {org_upper_name} GITHUB USERS', f'{org_users_count}'],
            [f'TOTAL {org_upper_name} GITHUB REPOSITORIES', f'{org_repos_count}'],
            [f'TOTAL {org_upper_name} USERS GITHUB REPOSITORIES', f'{org_users_repo_count}']
        ]
        print(tabulate(data, headers=['SCAN SUMMARY','DATA'], tablefmt="heavy_outline"))

    except Exception as error:
        log_error(f"Error in generating summery : {error}")
