# 使用官方的 Ubuntu 22.04 镜像作为基础镜像
FROM ubuntu:22.04

# 安装 wget、tar 和其他必要的工具
RUN apt-get update && apt-get install -y wget tar build-essential

# 定义环境变量，指定要下载的压缩文件的 URL
ENV DARKNET_URL "https://janction-origin-1324956105.cos.ap-seoul.myqcloud.com/darknet.tar.gz"
ENV JCT_URL "https://janction-origin-1324956105.cos.ap-seoul.myqcloud.com/jct-macos-amd64-0.0.9.tar.gz"

# 使用 wget 下载压缩文件
RUN wget $DARKNET_URL -O /tmp/darknet.tar.gz

# 创建 /app 目录
RUN mkdir -p /app

# 解压 darknet 压缩文件到 /app 目录
RUN tar -xzf /tmp/darknet.tar.gz -C /app

# 进入 darknet 文件夹
WORKDIR /app/darknet

# 编译 darknet
RUN make

# 回到 /app 目录
WORKDIR /app

# 下载 jct 压缩文件
RUN wget $JCT_URL -O /tmp/jct.tar.gz

# 解压 jct 压缩文件到 /app 目录
RUN tar -xzf /tmp/jct.tar.gz -C /app

# 赋予 jct 可执行权限
RUN chmod +x jct

# 定义容器启动时执行的命令
CMD ["./jct", "-config=config.json"]

# 可选：如果你需要在容器内部手动执行命令，可以移除上面的 CMD 指令
# 并在运行容器时附加到一个 shell，例如：
# docker run -it --name jct1 jct:0.1 /bin/bash





#################################################
## 使用官方的 Ubuntu 22.04 镜像作为基础镜像
#FROM ubuntu:22.04
#
## 安装 wget 和 tar 工具，用于下载和解压文件
#RUN apt-get update && apt-get install -y wget tar build-essential
#
## 定义环境变量，指定要下载的压缩文件的 URL
#ENV TAR_GZ_URL "https://janction-origin-1324956105.cos.ap-seoul.myqcloud.com/darknet.tar.gz"
#
## 使用 wget 下载压缩文件
#RUN wget $TAR_GZ_URL -O /tmp/darknet.tar.gz
#
#RUN mkdir -p /app
#
## 解压压缩文件到 /app 目录
#RUN tar -xzf /tmp/darknet.tar.gz -C /app
#
## 进入解压后的文件夹，这里假设文件夹名为 darknet
#WORKDIR /app/darknet
#
## 这里假设你需要执行的命令是 make，编译 darknet
## 请确保 makefile 存在于你的压缩文件中
#RUN make
#
## 可选：定义容器启动时执行的命令，例如启动一个服务
# CMD ["./darknet", "detector", "valid","cfg/voc.data","cfg/yolov3-voc.cfg","data/yolov3.weights"]