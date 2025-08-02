#!/usr/bin/env bash

# 用法: bash ./build/scripts/gen_server.sh

set -e

# 进入项目根目录（脚本在 build/scripts 下）
cd "$(dirname "$0")/../.." || exit 1

echo "🚀 正在从项目根目录执行 go run ./cmd/server/main.go gen api ... 进行 API 代码生成"

go run "./cmd/server/main.go" gen api \
  --config="./configs/config.yaml" \
  --env=".env" \
  --src="./api" \
  --output="./internal" \
  --log=false
