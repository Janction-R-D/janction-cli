FROM pytorch/pytorch:latest

RUN pip install numpy
RUN pip install torch torchvision

RUN apt update
RUN apt install -y curl wget jq vim

curl -s -o "$local_file" "https://janction-datas.s3.ap-northeast-1.amazonaws.com/tasks/$local_file"