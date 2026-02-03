#!/usr/bin/env bash

# 用法: bash ./build/scripts/start_client.sh

set -e

# 进入项目根目录（脚本在 build/scripts 下）
cd "$(dirname "$0")/../.." || exit 1

echo "🚀 正在从项目根目录执行 go run ./cmd/client/main.go ... 进行项目启动"

go run "./cmd/client/main.go" \
  --config="./configs/client.yaml" \
  --env=".env.client" \
  --show=false