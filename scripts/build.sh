#!/bin/bash
green() {
  "$@" | GREP_COLORS='mt=01;32' grep --color .
}

red() {
  "$@" | GREP_COLORS='mt=01;31' grep --color .
}

yellow() {
  "$@" | GREP_COLORS='mt=01;93' grep --color .
}

check_command_success() {
  CODE_TO_COMPARE_TO=$2
  RETURNED_CODE=$1
  if [ $RETURNED_CODE -ne $CODE_TO_COMPARE_TO ]; then
    if [[ $2 != "" ]]; then
      red echo "$3"
    fi
    exit 1
  fi
}

build_windows() {
  green echo "Starting building Hulotte for Windows"
  docker build -t hulotte_windows:build . -f ./build/package/Dockerfile.windows
  RESULT=$?
  check_command_success $RESULT 0 "Could not build Hulotte for windows"
  docker container create --name extract_hulotte hulotte_windows:build
  RESULT=$?
  check_command_success $RESULT 0 "Could not start builder container"
  docker container cp extract_hulotte:/go/src/github.com/clnbs/hulotte/hulotte_installer_windows.exe ./hulotte_installer_windows.exe
  RESULT=$?
  check_command_success $RESULT 0 "Could not extract binary from builder image"
  docker container rm -f extract_hulotte
  RESULT=$?
  check_command_success $RESULT 0 "Could not remove builder container"
}

clean_windows() {
  docker rmi hulotte_windows:build > /dev/null 2>&1
}

build_linux() {
  green echo "Starting building Hulotte for Linux"
  docker build -t hulotte_linux:build . -f ./build/package/Dockerfile.linux
  RESULT=$?
  check_command_success $RESULT 0 "Could not build Hulotte for Linux"
  docker container create --name extract_hulotte hulotte_linux:build
  RESULT=$?
  check_command_success $RESULT 0 "Could not start builder container"
  docker container cp extract_hulotte:/go/src/github.com/clnbs/hulotte/hulotte_installer_linux ./hulotte_installer_linux
  RESULT=$?
  check_command_success $RESULT 0 "Could not extract binary from builder image"
  docker container rm -f extract_hulotte
  RESULT=$?
  check_command_success $RESULT 0 "Could not remove builder container"
}

clean_linux() {
  docker rmi hulotte_linux:build > /dev/null 2>&1
}

build_linux