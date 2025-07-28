#!/bin/sh
set -e
(
  cd "$(dirname "$0")"
  go build -o /tmp/gel app/*.go
)
exec /tmp/gel "$@"
