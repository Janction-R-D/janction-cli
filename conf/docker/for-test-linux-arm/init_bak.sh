#!/bin/sh

# 确保 JCT_TASK 环境变量被设置
if [ -z "$JCT_TASK" ]; then
  echo "Error：JCT_TASK Not Found"
  exit 1
fi

# 下载配置文件
wget -P "/workspace" "https://janction-datas.s3.ap-northeast-1.amazonaws.com/linux-arm/config.json"

# 检查 wget 命令是否成功执行
if [ $? -ne 0 ]; then
  echo "Download Config Failed"
  exit 1
fi

# 下载任务脚本
wget -P "/workspace" "https://janction-datas.s3.ap-northeast-1.amazonaws.com/tasks/$JCT_TASK.py"

# 检查 wget 命令是否成功执行
if [ $? -ne 0 ]; then
  echo "Download Task Failed"
  exit 1
fi

# 运行 Janction Node
#/workspace/janction-node -config=/workspace/config.json