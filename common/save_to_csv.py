import pandas
from config.conf import *
from logger.log import *


def save_to_csv(filename):
    try:
        file_path = f"{OUTPUT_DIRECTORY}/{filename}.json"
        get_file = open(file_path, encoding="utf-8")
        read_file = pandas.read_json(get_file)
        read_file.to_csv(f"{OUTPUT_DIRECTORY}/{filename}.csv", encoding="utf-8")
        log_info(f"Saved data successfully into : {OUTPUT_DIRECTORY}/{filename}.csv")
    except Exception as error:
        log_error(f"error in converting JSON to CSV : {error}")
