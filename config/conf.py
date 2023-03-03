from dotenv import load_dotenv
from os import getenv
from cmd.cli import *
from logger.log import *


def validate_env_variable(variable_name):
    load_dotenv()
    get_variable = getenv(variable_name)

    if get_variable is None:
        log_error(f"{variable_name} is not configured in the environment variable")
        exit()
    else:
        log_info(f"{variable_name} is successfully configure in the environment variable")
        return get_variable


ORGANIZATION, OUTPUT_DIRECTORY, MONITOR_MODE, SLACK_NOTIFICATION = parse_cli_args()

ORG_USERS_FILE_NAME_DEFAULT = f"{ORGANIZATION}_users"
ORG_REPO_FILE_NAME_DEFAULT = f"{ORGANIZATION}_public_repositories"
ORG_USERS_REPO_FILE_NAME_DEFAULT = f"{ORGANIZATION}_users_public_repositories"

if not MONITOR_MODE:
    ORG_USERS_FILE_NAME = f"{ORGANIZATION}_users"
    ORG_REPO_FILE_NAME = f"{ORGANIZATION}_public_repositories"
    ORG_USERS_REPO_FILE_NAME = f"{ORGANIZATION}_users_public_repositories"
else:
    ORG_USERS_FILE_NAME = f"{ORGANIZATION}_users_tmp"
    ORG_REPO_FILE_NAME = f"{ORGANIZATION}_public_repositories_tmp"
    ORG_USERS_REPO_FILE_NAME = f"{ORGANIZATION}_users_public_repositories_tmp"

if SLACK_NOTIFICATION:
    SLACK_WEBHOOK = validate_env_variable("SLACK_WEBHOOK")

ORG_USER_FILE_PATH = f"{OUTPUT_DIRECTORY}/{ORG_USERS_FILE_NAME}.json"
ORG_REPO_FILE_PATH = f"{OUTPUT_DIRECTORY}/{ORG_REPO_FILE_NAME}.json"
ORG_USER_REPO_FILE_PATH = f"{OUTPUT_DIRECTORY}/{ORG_USERS_REPO_FILE_NAME}.json"

ORG_USER_FILE_PATH_DEFAULT = f"{OUTPUT_DIRECTORY}/{ORG_USERS_FILE_NAME_DEFAULT}.json"
ORG_REPO_FILE_PATH_DEFAULT = f"{OUTPUT_DIRECTORY}/{ORG_REPO_FILE_NAME_DEFAULT}.json"
ORG_USER_REPO_FILE_PATH_DEFAULT = f"{OUTPUT_DIRECTORY}/{ORG_USERS_REPO_FILE_NAME_DEFAULT}.json"

