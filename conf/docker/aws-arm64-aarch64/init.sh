#!/bin/sh
if [ -z "$JCT_TASK" ]; then
  echo "Error：JCT_TASK Not Found"
  exit 1
fi
wget -P /home/ubuntu -O /home/ubuntu/janction-node "https://janction-datas.s3.ap-northeast-1.amazonaws.com/linux-arm64-v8/janction-node"
if [ $? -ne 0 ]; then
  echo "Download janction-node Failed"
  exit 1
fi
chmod +x /home/ubuntu/janction-node
wget -P /home/ubuntu -O /home/ubuntu/config.json "https://janction-datas.s3.ap-northeast-1.amazonaws.com/linux-arm64-v8/config.json"
if [ $? -ne 0 ]; then
  echo "Download Config Failed"
  exit 1
fi
wget -P /home/ubuntu -O /home/ubuntu/$JCT_TASK.py "https://janction-datas.s3.ap-northeast-1.amazonaws.com/tasks/$JCT_TASK.py"
if [ $? -ne 0 ]; then
  echo "Download Task Failed"
  exit 1
fi
