import requests
from logger.log import *
from config.conf import *
def slack_notification(slack_message):
    try:
        if SLACK_NOTIFICATION:
            send_notification = requests.post(SLACK_WEBHOOK, json={"text": slack_message},
                                              headers={"Content-Type": "application/json"})
            slack_response = send_notification.text
            if slack_response == "ok":
                log_info("Slack message sent successfully")
            else:
                log_error("Error in sending slack message")
        else:
            return
    except Exception as error:
        log_error(f"Error in sending slack notification : {error}")

