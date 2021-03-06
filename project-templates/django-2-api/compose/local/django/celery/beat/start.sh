#!/usr/bin/env bash

set -o errexit
set -o pipefail
set -o nounset
set -o xtrace

rm -f './celerybeat.pid'
celery beat -A config --loglevel=DEBUG
