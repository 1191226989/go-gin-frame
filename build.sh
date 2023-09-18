#!/bin/bash

BUILD_DATE=$(date '+%Y%m%d')

# 构建docker image
docker build -f Dockerfile -t go-gin-frame:v${BUILD_DATE} .

# 推送docker hub