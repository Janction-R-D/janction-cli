FROM pytorch/pytorch:2.2.1-cuda12.1-cudnn8-runtime

RUN pip install numpy
RUN pip install torch torchvision

RUN apt update
RUN apt install -y wget jq vim

#ENV BIN_DOWNLOAD "https://janction-datas.s3.ap-northeast-1.amazonaws.com/linux-arm/janction-node"
#
#RUN wget $BIN_DOWNLOAD -O /workspace/janction-node
#RUN chmod +x /workspace/janction-node

COPY init.sh /usr/local/bin/init.sh

RUN chmod +x /usr/local/bin/init.sh


CMD ["sh", "-c", "/usr/local/bin/init.sh && /workspace/janction-node -config=/workspace/config.json"]

#CMD ["sh", "-c", "/usr/local/bin/init.sh && tail -f /dev/null"]

#
#CMD ["/workspace/janction-node", "-config=config.json"]
