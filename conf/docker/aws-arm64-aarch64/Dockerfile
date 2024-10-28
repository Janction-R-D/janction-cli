FROM armswdev/pytorch-arm-neoverse:r24.03-torch-2.2.1-onednn-acl

RUN pip install numpy
RUN pip install torch torchvision
RUN apt update
RUN apt install -y wget jq vim
COPY init.sh /usr/local/bin/init.sh
RUN chmod +x /usr/local/bin/init.sh
CMD ["sh", "-c", "/usr/local/bin/init.sh && /home/ubuntu/janction-node -config=/home/ubuntu/config.json"]

