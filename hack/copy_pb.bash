#!/bin/bash
set -eo pipefail

cd $(git rev-parse --show-toplevel)
../../hack/lib/copy_prebuilt_ts.bash
