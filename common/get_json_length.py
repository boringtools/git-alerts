import json
from logger.log import *

def get_json_len(file_path , json_attribute):
    try:
        file = open(file_path, "r")
        load_file = json.load(file)

        if not json_attribute:
            length = len(load_file)
            return length
        else:
            length = len(load_file[json_attribute])
            return length

    except json.decoder.JSONDecodeError as error:
        log_error("Error in reading JSON data")
