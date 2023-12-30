#! /bin/bash

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
scc
scc -f json | go run $SCRIPT_DIR/main.go --limit=$1 --readme-file=$2 --instachart-url=$3 --subtitle=$4@$5 --style=$6
