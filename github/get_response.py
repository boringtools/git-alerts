import requests
from github.github_config import *


def get_github_response(url):
    try:
        get_data = requests.get(url, params=github_request_params, headers=GITHUB_REQUEST_HEADERS)
        get_data_headers = get_data.headers
        get_data_response = get_data.json()

        pat_request_limit = get_data_headers["X-RateLimit-Remaining"]
        check_link_attr = get_data_headers.get("Link")

        if int(pat_request_limit) < 20:
            log_error("GitHub PAT request limit reached")
            exit()

        return get_data_headers, get_data_response, pat_request_limit, check_link_attr

    except requests.exceptions.RequestException as error:
        log_error(f"Failed to get response to the URL : {url}, ERROR: {error}  ")
