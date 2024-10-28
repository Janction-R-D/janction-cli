#!/bin/bash

# 发送 GET 请求并获取响应
#response=$(curl -s "http://localhost:8767/api/v1/task/list?task_type=cpu")
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
            echo "You chose $task"
            # 将 task 写入 docker 的 环境变量

            # 文件的本地路径和名称
#            local_file="/workspace/$task.py"
#
#            # 使用 curl 下载文件
#            curl -s -o "$local_file" "https://janction-datas.s3.ap-northeast-1.amazonaws.com/tasks/$local_file"
#
#            # 检查文件是否下载成功
#            if [ -f "$local_file" ]; then
#                echo "You chose successfully: $task"
#            else
#                echo "Failed to chose $task."
#            fi


            break
        else
            echo "Invalid option, please try again."
        fi
    done
else
    # 如果响应代码不是1000，输出错误信息
    msg=$(echo "$response" | jq '.msg')
    echo "Failed to fetch tasks: $msg"
fi