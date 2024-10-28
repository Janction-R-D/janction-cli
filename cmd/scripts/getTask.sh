#!/bin/bash

response=$(curl -s "https://www.janction.io/api/v1/task/list?task_type=cpu")

code=$(echo "$response" | jq '.code')
msg=$(echo "$response" | jq '.msg')
data=$(echo "$response" | jq '.data')

if [ $(echo "$response" | jq '.code') == "1000" ]; then
   echo "$response" | jq -c '.data[]' | while read item; do
     echo $item
       task_name=$(echo "$item" | jq -r '.task_name')
       task_code=$(echo "$item" | jq -r '.task_code')

       # 检查是否成功提取了 task_name 和 task_code
       if [ -z "$task_name" ] || [ -z "$task_code" ]; then
           echo "Error: Failed to extract task name or code."
           continue
       fi

       # 将 task_code 写入以 task_name 命名的文件中
       echo "$task_code" > "${task_name}.py"
       echo "Task code for $task_name has been written to ${task_name}.py"
    done
else
    echo "Failed to fetch tasks: $msg"
fi