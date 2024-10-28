#!/bin/bash

#mkdir /workspace

# 读取环境变量
cd /workspace && wget -P "/workspace" "https://janction-datas.s3.ap-northeast-1.amazonaws.com/tasks/$JCT_TASK.py"

echo $JCT_TASK

echo "[--$JCT_TASK--]" >> eee3



#1. start.sh 脚本用来选择任务 将环境变量设置进 docker 容器
#
#
#-e JCT_CPU=Cortex-A55 \
#-e JCT_GPU=none \
#-e JCT_GPU_ID=none \
#-e JCT_TASK=cpu_simple_linear_regression \
#-e JCT_TASK_TYPE=cpu \
#-e JCT_USE_DEVICE=cpu \
#-e PRIVATE_KEY=0x4a6345abcdef17e36ba4a6b4d238ff944bacb478cbed5efca12bbc64a6a9cbb3
#
#2. 在 init.sh 中下载任务脚本
#
#3. go二进制读环境变量的任务来跑任务
#
#4. 跑的任务注意一下成功失败的标志位判断。先打印出来日志