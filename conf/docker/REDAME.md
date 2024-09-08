

## 构建命令 
docker build -t jct:0.0.9 .
docker build -t jct-windows:0.0.9 .
docker build -t jct-linux-amd64:0.0.9 .
docker build -t jct-linux-arm:0.0.9 .
docker build -t jct-macos-arm:0.0.9 .
docker build -t jct-macos-amd64:0.0.9 .
## 运行命令

docker run -d -e PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --name janction-node1 roddyneo/jct-linux-amd64:0.0.9
docker run -d -e PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --name janction-node2 roddyneo/jct-linux-arm:0.0.9
docker run -d -e PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --name janction-node3 roddyneo/jct-macos-amd64:0.0.9
docker run -d -e PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --name janction-node4 roddyneo/jct-macos-arm:0.0.9
docker run -d -e PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --name janction-node5 roddyneo/jct-windows-amd64:0.0.9

## 调试命令
docker run -it --name jct2 jct:0.0.8 /bin/bash

## 编译
CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -ldflags='-s -w -extldflags "-static -fpic"' -o jct  debug.go
CGO_ENABLE=0 GOOS=linux GOARCH=amd64 go build -ldflags='-s -w -extldflags "-static -fpic"' -o jct  debug.go

docker run -d --name jct1 roddyneo/jct:latest -e PRIVATE_KEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

https://hub.docker.com/repositories/roddyneo