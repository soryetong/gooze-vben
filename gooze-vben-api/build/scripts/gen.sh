#!/usr/bin/env bash

# 用法: sh ./build/scripts/gen.sh <mainPath> <configPath> <envPath> <srcPath> <outputPath> <needLog>
# 示例: sh ./build/scripts/gen.sh ./cmd/admin/main.go ./configs/admin.yaml ./env.admin ./api ./internal true

set -e

# 进入项目根目录（脚本在 build/scripts 下）
cd "$(dirname "$0")/../.." || exit 1

MAIN_PATH="$1"
CONFIG_PATH="$2"
ENV_PATH="$3"
SRC_PATH="$4"
OUTPUT_PATH="$5"
NEED_LOG="$6"

echo "🚀 正在从项目根目录执行 go run $MAIN_PATH gen api ... 进行 API 代码生成"

go run "$MAIN_PATH" gen api \
  --config="$CONFIG_PATH" \
  --env="$ENV_PATH" \
  --src="$SRC_PATH" \
  --output="$OUTPUT_PATH" \
  --log=$6