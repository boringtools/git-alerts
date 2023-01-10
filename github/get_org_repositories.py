from github.get_response import *


def get_org_repositories():
    log_info(f"fetching {ORGANIZATION} public repositories")

    try:
        headers, response, pat_limit, link_attr = get_github_response(f"{github_url_org}/repos")
        log_info(f"Remaining request limit: {pat_limit}")

        if link_attr is None:
            for repos in response:
                visibility = repos["visibility"]
                if visibility == "public":
                    organization_directory["repository"].append(repos["html_url"])
                    organization_directory["fork"].append(repos["fork"])
                    organization_directory["created"].append(repos["created_at"])
                    organization_directory["updated"].append(repos["updated_at"])
                    organization_directory["pushed"].append(repos["pushed_at"])
        else:
            page_length = int(link_attr.split(",")[1].split("page=")[1].split("&")[0])

            for pages in range(1, page_length + 1):
                github_request_params["page"] = pages
                headers, response, pat_limit, link_attr = get_github_response(f"{github_url_org}/repos")

                for repos in response:
                    visibility = repos["visibility"]
                    if visibility == "public":
                        organization_directory["repository"].append(repos["html_url"])
                        organization_directory["fork"].append(repos["fork"])
                        organization_directory["created"].append(repos["created_at"])
                        organization_directory["updated"].append(repos["updated_at"])
                        organization_directory["pushed"].append(repos["pushed_at"])

        return organization_directory

    except requests.exceptions.RequestException as error:
        log_error(f"Error in fetching {ORGANIZATION} repositories : {error}")
        exit()
