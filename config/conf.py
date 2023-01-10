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


ORGANIZATION, OUTPUT_DIRECTORY = parse_cli_args()

ORG_USERS_FILE_NAME = f"{ORGANIZATION}_users"
ORG_REPO_FILE_NAME = f"{ORGANIZATION}_public_repositories"
ORG_USERS_REPO_FILE_NAME = f"{ORGANIZATION}_users_public_repositories"



