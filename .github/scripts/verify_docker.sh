#!/bin/bash

set -euo pipefail

# verify_docker.sh invokes the given Docker image to run `waypoint version` and inspect its output.
# If its output doesn't match the version given, the script will exit 1 and report why it failed.
# This is meant to be run as part of the build workflow to verify the built image meets some basic
# criteria for validity.
#
# Because this is meant to be run as the `smoke_test` for the docker-build workflow, the script expects
# the image name parameter to be provided by the `IMAGE_NAME` environment variable, rather than a
# positional argument.

function usage {
  echo "IMAGE_NAME=<image uri> ./verify_docker.sh <expect_version>"
}

function main {
  local image_name="${IMAGE_NAME:-}"
  local expect_version="${1:-}"
  local got_version

  if [[ -z "${image_name}" ]]; then
    echo "ERROR: IMAGE_NAME is not set"
    usage
    exit 1
  fi

  if [[ -z "${expect_version}" ]]; then
    echo "ERROR: expected version argument is required"
    usage
    exit 1
  fi

  full_version=$(docker run --rm "${image_name}" version)
  echo "Full version: ${full_version}"
  got_version="$( awk '{print $2}' <(head -n1 <(docker run --rm "${image_name}" version)) )"
  if [[ "${got_version}" != "v${expect_version}" ]]; then
    echo "Version Test FAILED"
    echo "Got: ${got_version}, Want: v${expect_version}"
    exit 1
  fi
  echo "Test PASSED"
}

main "$@"
