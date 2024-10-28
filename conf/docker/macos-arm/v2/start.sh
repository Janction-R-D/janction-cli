#!/bin/bash

container=$(docker ps -aq -f name=janction-node)

if [ -n "$container" ]; then
    # 如果容器存在，停止并删除它
    echo "Stopping and removing the janction-node container..."
    docker stop janction-node
    docker rm janction-node
    echo "The janction-node container has been stopped and removed."
fi

has_gpu=false

JCT_CPU=$(lscpu | grep "Model name" | awk -F':' '{print $2}' | xargs echo | sed 's/^[ \t]*//')
JCT_GPU="none"
JCT_GPU_ID="none"
JCT_TASK_TYPE="cpu"
JCT_USE_DEVICE="cpu"

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
    JCT_GPU=$(nvidia-smi -L | head -n 1 |awk -F':' '{print $2}'| awk -F'(' '{print $1}'| xargs echo | sed 's/^[ \t]*//')
    JCT_GPU_ID=$(nvidia-smi -L | head -n 1 | grep -o '\GPU-[0-9a-f-]*' | xargs echo | sed 's/^[ \t]*//')
    has_gpu=true
    nvidia-smi
else
    echo "[WARN] No GPU driver detected. Please install the GPU driver first."
    echo "You can only run CPU tasks"
fi

validate_private_key() {
    local private_key="$1"
    if [[ $private_key =~ ^0x[a-fA-F0-9]{64}$ ]]; then
        return 0
    else
        return 1
    fi
}

while true; do
    read -p "Please enter your private key: " PRIVATE_KEY
    if validate_private_key "$PRIVATE_KEY"; then
        echo "You entered the following private key:"
        echo "PRIVATE=$PRIVATE_KEY"
        break
    else
        echo "Error: The private key must start with '0x' and be 66 characters long."
        echo "Please try again."
    fi
done

if [ "$has_gpu" = true ]; then
    echo "Select the task type:"
    select task_type in "GPU" "CPU"; do
        case $task_type in
            "GPU" )
                response=$(curl -s "https://www.janction.io/api/v1/task/list?task_type=gpu")
                JCT_TASK_TYPE="gpu"
                JCT_USE_DEVICE="gpu"
                break
                ;;
            "CPU" )
                response=$(curl -s "https://www.janction.io/api/v1/task/list?task_type=cpu")
                JCT_TASK_TYPE="cpu"
                JCT_USE_DEVICE="cpu"
                break
                ;;
            * ) echo "Invalid option. Please choose again." ;;
        esac
    done
else
    response=$(curl -s "https://www.janction.io/api/v1/task/list?task_type=cpu")
fi

task_names=()
if [ $(echo "$response" | jq '.code') == "1000" ]; then
    while IFS= read -r line; do
        task_names+=("$line")
    done < <(echo "$response" | jq -r '.data[].task_name')
    echo task_names
    echo "Available tasks:"
    select task in "${task_names[@]}"; do
        if [ -n "$task" ]; then
            echo "You choose $task, Starting Janction Node"
            break
        else
            echo "Invalid"
        fi
    done

    if docker run -d -e JCT_CPU="$JCT_CPU" -e JCT_GPU="$JCT_GPU" -e JCT_GPU_ID="$JCT_GPU_ID" -e JCT_TASK="$task" -e JCT_TASK_TYPE=$JCT_TASK_TYPE -e JCT_USE_DEVICE=$JCT_USE_DEVICE -e PRIVATE_KEY=$PRIVATE_KEY  --name janction-node roddyneo/janction-node-mac-arm:0.2.0; then
        echo "Janction node started successfully."
    else
        echo "Failed to start Janction node."
        exit 1
    fi


else
    msg=$(echo "$response" | jq '.msg')
    echo "Failed to fetch tasks: $msg"
fi