FROM pytorch/pytorch:2.2.1-cuda12.1-cudnn8-runtime

RUN pip install numpy
RUN pip install torch torchvision
RUN apt update
RUN apt install -y wget jq vim
COPY init.sh /usr/local/bin/init.sh
RUN chmod +x /usr/local/bin/init.sh
CMD ["sh", "-c", "/usr/local/bin/init.sh && /workspace/janction-node -config=/workspace/config.json"]

