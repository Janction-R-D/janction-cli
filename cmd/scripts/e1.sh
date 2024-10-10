#!/bin/bash

# 定义一个函数来验证私钥格式
validate_private_key() {
    local private_key="$1"
    # 检查私钥是否以 "0x" 开头并且长度为66
    if [[ $private_key =~ ^0x[a-fA-F0-9]{64}$ ]]; then
        return 0 # 格式正确
    else
        return 1 # 格式错误
    fi
}

# 提示用户输入私钥
while true; do
    read -p "Please enter your private key: " PRIVATE

    # 验证私钥格式
    if validate_private_key "$PRIVATE"; then
        echo "You entered the following private key:"
        echo "PRIVATE=$PRIVATE"
        break # 退出循环
    else
        echo "Error: The private key must start with '0x' and be 66 characters long."
        echo "Please try again."
    fi
done