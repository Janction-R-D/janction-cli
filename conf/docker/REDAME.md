

## 构建命令 
docker build -t jct:0.4 .
## 运行命令
docker run -it --name jct1 jct:0.1 
## 调试命令
docker run -it --name jct2 jct:0.4 /bin/bash