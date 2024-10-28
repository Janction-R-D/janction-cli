#!/bin/bash

echo "Current operating system information:"
uname -a

echo "Check Docker:"
if which docker > /dev/null; then
    echo "Docker is installed."
    docker --version
else
    echo "Docker is not installed."
    curl https://get.docker.com/ | sh
fi

echo "Check if there is a graphics card driver (NVIDIA driver):"
if which nvidia-smi > /dev/null; then
    echo "NVIDIA driver installed."
    nvidia-smi
else
    echo "[ERROR] No GPU driver detected. Please install the GPU driver first."
fi