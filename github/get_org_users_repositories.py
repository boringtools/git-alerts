import json

from github.get_response import *


def get_org_users_repositories():
    log_info(f"Fetching {ORGANIZATION} users public repositories")

    try:
        org_users = open(f"{OUTPUT_DIRECTORY}/{ORG_USERS_FILE_NAME}.json", "r")
        load_org_users = json.load(org_users)
        users = load_org_users["usernames"]

        for user in users:
            headers, response, pat_limit, link_attr = get_github_response(f"{github_url_user}/{user}/repos")
            if not link_attr:
                for repo in response:
                    organization_users_directory["repository"].append(repo["html_url"])
                    organization_users_directory["fork"].append(repo["fork"])
                    organization_users_directory["created"].append(repo["created_at"])
                    organization_users_directory["updated"].append(repo["updated_at"])
                    organization_users_directory["pushed"].append(repo["pushed_at"])
            else:
                page_length = int(link_attr.split(",")[1].split("page=")[1].split("&")[0])
                for pages in range(1, page_length + 1):
                    github_request_params["page"] = pages
                    headers_2, response_2, pat_limit_2, link_attr_2 = get_github_response(f"{github_url_user}/{user}/repos")
                    for repos in response_2:
                        organization_users_directory["repository"].append(repos["html_url"])
                        organization_users_directory["fork"].append(repos["fork"])
                        organization_users_directory["created"].append(repos["created_at"])
                        organization_users_directory["updated"].append(repos["updated_at"])
                        organization_users_directory["pushed"].append(repos["pushed_at"])

        log_info(f"Remaining request limit: {pat_limit}")
        return organization_users_directory

    except json.decoder.JSONDecodeError as error:
        log_error(f"Error in reading JSON data : {error}")
        exit()
