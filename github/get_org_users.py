from github.get_response import *


def get_org_users():
    log_info(f"Fetching {ORGANIZATION} users ")

    try:
        headers, response, pat_limit, link_attr  = get_github_response(f"{github_url_org}/members")
        log_info(f"Remaining request limit: {pat_limit}")

        if link_attr is None:

            for users in response:
                username = users["login"]
                users_directory["usernames"].append(username)
        else:
            page_length = int(link_attr.split(",")[1].split("page=")[1].split("&")[0])

            for pages in range(1, page_length + 1):
                github_request_params["page"] = pages
                headers, response, pat_limit, link_attr = get_github_response(f"{github_url_org}/members")

                for users in response:
                    username = users["login"]
                    users_directory["usernames"].append(username)
        return users_directory

    except requests.exceptions.RequestException as error:
        log_error(f"Error in fetching {ORGANIZATION} users : {error}")
        exit()
