操作：
先在 aarch64服务器上
先编译 pytorch 镜像


修改内容：

linux arm  / mac arm 是一样的 node bin
这里对应 aws 上 linux-arm

init的脚本 : 修改 linux-arm
每次在启动的时候：

start .sh 修改





# GOARCH=arm GOOS=linux go build -v -ldflags "-w -s" -o janction-node .

最终

docker build --platform linux/arm64 -t janction-node-arm64-v8:0.2.0 .


docker tag 1cc33dfb8354fc7e52bd80e62b85b043f226649ae109f867d43d2608b53fe583 roddyneo/janction-node-arm64-v8:0.2.0

docker push roddyneo/janction-node-arm64-v8:0.2.0

docker pull roddyneo/janction-node-arm64-v8:0.2.0

