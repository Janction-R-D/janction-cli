

## 构建命令 
docker build -t jct:0.0.8 .
## 运行命令
docker run -d --name jct1 jct:0.0.8  -e private_key=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
## 调试命令
docker run -it --name jct2 jct:0.0.8 /bin/bash

## 编译
CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -ldflags='-s -w -extldflags "-static -fpic"' -o jct  main.go

docker run -d --name jct1 roddyneo/jct:latest -e private_key=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

