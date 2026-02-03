#!/usr/bin/env bash

# 用法: bash ./build/scripts/gen_client.sh

set -e

# 进入项目根目录（脚本在 build/scripts 下）
cd "$(dirname "$0")/../.." || exit 1

echo "🚀 正在从项目根目录执行 go run ./cmd/client/main.go gen api ... 进行 API 代码生成"

go run "./cmd/client/main.go" gen api \
  --config="./configs/client.yaml" \
  --env=".env.client" \
  --src="./api/client" \
  --output="./internal" \
  --log=false
