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

docker build --platform linux/arm64/v8 -t janction-node-arm64-v8:0.2.0 .


docker tag 7445432fbc48 roddyneo/janction-node-arm64-v8:0.2.0

docker push roddyneo/janction-node-arm64-v8:0.2.0

docker pull roddyneo/janction-node-arm64-v8:0.2.0





docker run --rm -it -e JCT_CPU="Neoverse-V1 AWS Graviton3 AWS Graviton3 CPU @ 2.6GHz" -e JCT_GPU=none -e JCT_GPU_ID=none -e JCT_TASK=cpu_simple_linear_regression -e JCT_TASK_TYPE=cpu -e JCT_USE_DEVICE=cpu -e PRIVATE_KEY=0x4a6345abcdef17e36ba4a6b4d238ff944bacb478cbed5efca12bbc64a6a9cbb3  roddyneo/janction-node-arm64-v8:0.2.0


sh链接：
https://ap-northeast-1.console.aws.amazon.com/s3/object/janction-datas?region=ap-northeast-1&bucketType=general&prefix=linux-arm64-v8/start.sh