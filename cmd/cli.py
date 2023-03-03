import argparse


def parse_cli_args():
    parser = argparse.ArgumentParser(
        description="GitAlert Tool"
    )

    parser.add_argument("-o", "--org", required=True)
    parser.add_argument("-d", "--output-directory", required=False, default="/tmp")
    parser.add_argument("-m", "--monitor-mode", required=False, default=False)
    parser.add_argument("-s", "--slack-notification", required=False, default=False)
    args = parser.parse_args()

    organization = args.org
    output_directory = args.output_directory
    monitor_mode = args.monitor_mode
    slack_notification = args.slack_notification
    return organization, output_directory, monitor_mode, slack_notification
