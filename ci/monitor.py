import shutil

from tabulate import tabulate
from ci.get_new_repos import *
from notification.slack import *

def  monitor():

    if MONITOR_MODE:
        new_repositories = get_new_repos()

        if len(new_repositories):
            message = ":eyes: *New Public Repositories*"
            data = []

            for repos in new_repositories:
                message = message + "\n :point_right: "+repos
                data.append([repos])

            print(tabulate(data, headers=['New Public Repositories'], tablefmt="heavy_outline"))
            slack_notification(message)

            shutil.move(ORG_USER_FILE_PATH, ORG_USER_FILE_PATH_DEFAULT)
            shutil.move(ORG_REPO_FILE_PATH, ORG_REPO_FILE_PATH_DEFAULT)
            shutil.move(ORG_USER_REPO_FILE_PATH, ORG_USER_REPO_FILE_PATH_DEFAULT)
        else:
            log_info("No new public repository detected")
    else:
        return
