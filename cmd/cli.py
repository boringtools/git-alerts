import argparse


def parse_cli_args():
    parser = argparse.ArgumentParser(
        description="GitAlert Tool"
    )

    parser.add_argument("-o", "--org", required=True)
    parser.add_argument("-d", "--output-directory", required=False, default="/tmp")
    args = parser.parse_args()

    organization = args.org
    output_directory = args.output_directory
    return organization, output_directory
