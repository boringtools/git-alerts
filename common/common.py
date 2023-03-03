import json
import os.path
from logger.log import *

def check_file_exists(file_path):
    if os.path.exists(file_path):
        return True
    else:
        return False

def read_json_file(file_path, json_attribute):
    try:
        if check_file_exists(file_path):
            file = open(file_path, "r")
            load_file = json.load(file)

            if not json_attribute:
                return load_file
            else:
                return load_file[json_attribute]

    except json.decoder.JSONDecodeError as error:
        log_error("Error in reading JSON data")
