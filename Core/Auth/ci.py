import os
import re
import sys
import semver
import subprocess
import logging
import argparse

# Set up logging
logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)


def git(*args):
    """Run a git command and return its output."""
    try:
        return subprocess.check_output(["git"] + list(args))
    except subprocess.CalledProcessError as e:
        logger.error(f"Git command failed: {e}")
        sys.exit(1)


def tag_repo(tag):
    """Tag the git repository and push the tag to origin."""
    url = os.environ.get("CI_REPOSITORY_URL")
    if not url:
        logger.error("CI_REPOSITORY_URL environment variable not set.")
        sys.exit(1)

    push_url = re.sub(r'.+@([^/]+)/', r'git@\1:', url)
    git("remote", "set-url", "--push", "origin", push_url)
    git("tag", tag)
    git("push", "origin", tag)


def bump(version, part):
    """Bump the specified part of a semantic version string."""
    if part == "major":
        return semver.bump_major(version)
    elif part == "minor":
        return semver.bump_minor(version)
    elif part == "patch":
        return semver.bump_patch(version)
    else:
        logger.error(f"Invalid version part: {part}")
        sys.exit(1)


def main(part):
    """The main script function."""
    try:
        latest = git("describe", "--tags").decode().strip()
    except subprocess.CalledProcessError:
        version = "1.0.0"
    else:
        if '-' not in latest:
            print(latest)
            return 0
        version = bump(latest, part)

    tag_repo(version)
    print(version)
    return 0


if __name__ == "__main__":
    parser = argparse.ArgumentParser()
    parser.add_argument('part', choices=['major', 'minor', 'patch'], help='The part of the version to bump.')
    args = parser.parse_args()
    sys.exit(main(args.part))
