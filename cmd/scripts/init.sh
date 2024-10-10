#!/bin/bash

# 发送 GET 请求并获取响应
response=$(curl -s "http://localhost:8767/api/v1/task/list?task_type=cpu")


# 解析 JSON 数据
code=$(echo "$response" | jq '.code')
msg=$(echo "$response" | jq '.msg')
data=$(echo "$response" | jq '.data')

echo "$msg"
echo "$data"

if [ "$code" == "1000" ] ; then
  echo "123"
fi

# 检查响应代码和消息
if [ "$code" == "1000" ]; then
    # 将任务名称提取到数组中
    mapfile -t tasks <<< "$data"

    # 显示任务列表并让用户选择
    echo "Available tasks:"
    select task in "${tasks[@]}"; do
        if [ -n "$task" ]; then
            # 用户选择后，输出选择的任务
            echo "You chose $task"
            break
        else
            echo "Invalid option, please try again."
        fi
    done
else
    echo "Failed to fetch tasks: $msg"
fi