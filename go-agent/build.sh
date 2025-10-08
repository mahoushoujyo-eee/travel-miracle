#!/bin/bash
RUN_NAME=hertz_service
echo "构建 Travel Planner Go Agent..."
go build -o ${RUN_NAME}
echo "构建完成: ${RUN_NAME}"