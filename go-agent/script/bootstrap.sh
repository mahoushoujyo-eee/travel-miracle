#!/bin/bash

# 获取脚本所在目录
CURDIR=$(cd $(dirname $0); pwd)
PROJECT_ROOT=$(cd $CURDIR/..; pwd)
BinaryName=hertz_service

echo "=== Travel Planner Go Agent 构建和执行脚本 ==="
echo "项目根目录: $PROJECT_ROOT"

# 切换到项目根目录
cd $PROJECT_ROOT

# 检查依赖
echo "检查Go模块依赖..."
go mod tidy

# 构建项目
echo "编译Go程序..."
go build -o ${BinaryName}

if [ $? -eq 0 ]; then
    echo "构建成功！"
else
    echo "构建失败！"
    exit 1
fi

# 执行程序
echo "启动 Travel Planner Go Agent..."
echo "=================================="

exec ./${BinaryName}