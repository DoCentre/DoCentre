#!/usr/bin/env sh

set -eux

go install github.com/githubnemo/CompileDaemon@latest &&
    CompileDaemon -command="./docentre" -exclude-dir=".*"
