FROM docker.io/bitnami/minideb:bookworm

ARG TARGETARCH

ENV HOME="/" \
    OS_ARCH="${TARGETARCH:-amd64}" \
    OS_FLAVOUR="debian-12" \
    OS_NAME="linux"

# 安装 wget、tar 和其他必要的工具
RUN apt-get update && apt-get install -y curl wget tar build-essential


