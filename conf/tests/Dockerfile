FROM bitnami/minideb:bookworm

RUN apt-get update && apt-get install -y curl wget tar build-essential

COPY init.sh /usr/local/bin/init.sh

RUN chmod +x /usr/local/bin/init.sh

CMD ["sh", "-c", "/usr/local/bin/init.sh && tail -f /dev/null"]