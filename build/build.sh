#!/bin/sh

# execute by Jenkins docker

export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on

SCRIPT_DIR="."
PROJECT_DIR="${SCRIPT_DIR}/"
RES_FILE_NAME="plugin-based-executor"
RES_FILE_PATH="${PROJECT_DIR}/${RES_FILE_NAME}"

# build_app 编译
build_app() {
    echo "enter download_dependency"
    if [ ! -d "${PROJECT_DIR}" ]; then
        echo "do not have project dir"
        return 1
    fi

    go mod download
    CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static" -w -s' -o "${RES_FILE_PATH}" ./src/main

    chmod +x "${RES_FILE_PATH}"
    return 0
}

# main
main() {
    echo "remove old"
    rm -rf "${RES_FILE_PATH}"

    build_app
    build_res=$?
    if [ "$build_res" != 0 ]; then
        echo "build error"
    fi
    echo "build success"
}

main "$@"