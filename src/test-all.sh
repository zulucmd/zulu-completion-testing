#!/usr/bin/env bash

set -euo pipefail

# This script runs completion tests in different environments and different shells.

# Get path to docker or podman binary
CONTAINER_ENGINE="$(command -v podman docker | head -n1)"

if [ -z "$CONTAINER_ENGINE" ]; then
  echo "Missing 'docker' or 'podman' which is required for these tests"
  exit 2
fi

BASE_DIR=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")/.." &>/dev/null && pwd)

export TESTS_DIR="${BASE_DIR}/tests"
export TESTPROG_DIR="${BASE_DIR}/testprog"
export TESTING_DIR="${BASE_DIR}/testingdir"

# Run all tests, even if there is a failure.
# But remember if there was any failure to report it at the end.
set +e
GOT_FAILURE=0
trap "GOT_FAILURE=1" ERR

declare -A test_cases=()

for SHELL_TYPE in "$@"; do
  mapfile -d '' tests < <(find "${BASE_DIR}/tests" -name "Dockerfile.*-${SHELL_TYPE}-*" -print0)
  for testFile in "${tests[@]}"; do
    filename="$(basename "$testFile")"
    testName="${filename:11}" # strip Dockerfile.

    imageName="comp-test:$SHELL_TYPE-$testName"
    test_cases[$imageName]="$SHELL_TYPE"

    (
      exec > >(trap "" INT TERM; sed 's/^/'"$testName"': /')
      exec 2> >(trap "" INT TERM; sed 's/^/'"$testName"': /' >&2)
      $CONTAINER_ENGINE build -t "${imageName}" "${BASE_DIR}" -f "$testFile"
    ) &
  done
done

wait

for imageName in "${!test_cases[@]}"; do
  shellType=${test_cases[$imageName]}
  "$CONTAINER_ENGINE" run --rm "${imageName}" "tests/comp-tests.$shellType"
done
# Indicate if anything failed during the run
exit ${GOT_FAILURE}
