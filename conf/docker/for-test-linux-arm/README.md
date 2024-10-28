docker build --platform=linux/arm64,linux/arm64/v8,linux/amd64 -t janction-node-linux:0.2.0 .


docker tag ea53b55367f84a732f269dea70288e0daef04bda505a084ab467a90718ebd2ca roddyneo/janction-node-linux:0.2.0

docker push roddyneo/janction-node-linux:0.2.0

docker pull roddyneo/janction-node-linux:0.2.0




docker run --rm -it -e JCT_CPU="Neoverse-V1 AWS Graviton3 AWS Graviton3 CPU @ 2.6GHz" -e JCT_GPU=none -e JCT_GPU_ID=none -e JCT_TASK=cpu_simple_linear_regression -e JCT_TASK_TYPE=cpu -e JCT_USE_DEVICE=cpu -e PRIVATE_KEY=0x4a6345abcdef17e36ba4a6b4d238ff944bacb478cbed5efca12bbc64a6a9cbb3  roddyneo/janction-node-linux:0.2.0