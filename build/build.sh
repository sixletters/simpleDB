#!/bin/bash

WORK_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd -P)"

cd "$WORK_DIR" || exit 1

LINUX_OUTPUT_DIR="$WORK_DIR/output/bin/"

mkdir -p "$LINUX_OUTPUT_DIR"

BUILD_DIR="$WORK_DIR/cmd"
cd "$BUILD_DIR" || exit 1

export GO111MODULE=on
env go build -o "$LINUX_OUTPUT_DIR" .
