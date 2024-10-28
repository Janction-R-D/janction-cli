
包：
https://janction-datas.s3.ap-northeast-1.amazonaws.com/linux-arm/jct-node-cli

init的脚本
每次在启动的时候：
下载文件

$local_file 取环境变量
curl -s -o "$local_file" "https://janction-datas.s3.ap-northeast-1.amazonaws.com/tasks/$local_file"



init.sh : 修改 mac-arm



# GOARCH=arm GOOS=linux go build -v -ldflags "-w -s" -o janction-node .

最终

docker build --platform linux/arm64 -t janction-node-mac-arm:0.2.0 .


docker tag 81c3a506b8b8acf3554aba19e171d669e2307040b67f7b3931bba0d47e1a9603 roddyneo/janction-node-mac-arm:0.2.0

docker push roddyneo/janction-node-mac-arm:0.2.0

docker pull roddyneo/janction-node-mac-arm:0.2.0



test:
zai