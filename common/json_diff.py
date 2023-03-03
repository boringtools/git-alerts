import json
from logger.log import *
from common.common import *
from config.conf import *


def get_json_diff(new_file, old_file, json_attribute):
    if check_file_exists(new_file) and check_file_exists(old_file):
        updated_file = read_json_file(new_file, json_attribute)
        old_file = read_json_file(old_file, json_attribute)
        new_added_repos = list(set(updated_file) - set(old_file))

        if new_added_repos:
            return new_added_repos
        else:
            return



