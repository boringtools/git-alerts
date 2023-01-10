import json

from config.conf import *
from logger.log import *


def save_to_json(data, filename):
    try:
        file_path = f"{OUTPUT_DIRECTORY}/{filename}.json"
        save_data = open(file_path, "w")
        json.dump(data, save_data)
        save_data.close()
        log_info(f"Saved data successfully into : {file_path}")
    except json.decoder.JSONDecodeError as error:
        log_error(f"Error in saving JSON data : {error}")