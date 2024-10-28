#!/bin/bash

echo "Current operating system information:"
uname -a

echo "Check Docker:"
if which docker > /dev/null; then
    echo "Docker is installed."
    docker --version
else
    echo "[ERROR] Docker is not installed."
    exit 1
fi

echo "Check if there is a graphics card driver (NVIDIA driver):"
if which nvidia-smi > /dev/null; then
    echo "NVIDIA driver installed."
    nvidia-smi
else
    echo "[WARN] No GPU driver detected. Please install the GPU driver first."
    echo "You can only run GPU tasks"
fi

# 发送 GET 请求并获取响应
response=$(curl -s "https://www.janction.io/api/v1/task/list?task_type=cpu")

# 检查响应代码是否为1000
if [ $(echo "$response" | jq '.code') == "1000" ]; then
    # 从 JSON 响应中提取 task_name 到数组
    mapfile -t task_names < <(echo "$response" | jq -r '.data[].task_name')

    # 显示任务列表并让用户选择
    echo "Available tasks:"
    select task in "${task_names[@]}"; do
        if [ -n "$task" ]; then
            # 用户选择后，输出选择的任务
            echo "You chose $task, Starting the Janction node"
            break
        else
            echo "Invalid option, please try again."
        fi
    done
    # 执行 docker run 命令并检查结果
    if docker run -d -e PRIVATE_KEY=0xaaa -e TASK="$task" --name janction-node roddyneo/jct-linux-amd64:0.0.9; then
        echo "Janction node started successfully."
    else
        echo "Failed to start Janction node."
        exit 1
    fi
else
    # 如果响应代码不是1000，输出错误信息
    msg=$(echo "$response" | jq '.msg')
    echo "Failed to fetch tasks: $msg"
fi