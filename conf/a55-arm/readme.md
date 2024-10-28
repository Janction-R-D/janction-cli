
包：
https://janction-datas.s3.ap-northeast-1.amazonaws.com/linux-arm/jct-node-cli

init的脚本
每次在启动的时候：
下载文件

$local_file 取环境变量
curl -s -o "$local_file" "https://janction-datas.s3.ap-northeast-1.amazonaws.com/tasks/$local_file"





# GOARCH=arm GOOS=linux go build -v -ldflags "-w -s" -o janction-node .

最终

docker build --platform linux/arm -t janction-node-arm:0.1.2 .


docker tag e92ed6a2eaf070e0f6ff57ba405af2975ee1a70c848ebf7224c95b8e9a3322a4 roddyneo/janction-node-arm:0.1.2

docker push roddyneo/janction-node-arm:0.1.2



docker pull roddyneo/janction-node-arm:0.1.2



docker run -d -e JCT_CPU="Neoverse-V1 AWS Graviton3 AWS Graviton3 CPU @ 2.6GHz" -e JCT_GPU=none -e JCT_GPU_ID=none -e JCT_TASK=cpu_simple_linear_regression -e JCT_TASK_TYPE=cpu -e JCT_USE_DEVICE=cpu -e PRIVATE_KEY=0x4a6345abcdef17e36ba4a6b4d238ff944bacb478cbed5efca12bbc64a6a9cbb3 --name janction-node roddyneo/janction-node-arm:0.1.2
