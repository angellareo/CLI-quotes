#!/bin/bash

# Default language
LANGUAGE="en"

# Parse arguments
FILES=()
while [[ $# -gt 0 ]]; do
  case $1 in
    -l|--lang)
      LANGUAGE="$2"
      shift 2
      ;;
    *)
      FILES+=("$1")
      shift
      ;;
  esac
done

# Copy json to go source code for them to be
# hardcoded in the binary
node ./build-cli.js "$LANGUAGE" "${FILES[@]}"

# build the go
cd cli
go fmt
go build
