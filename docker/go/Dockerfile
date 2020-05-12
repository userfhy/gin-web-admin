FROM golang:latest

ENV TIME_ZONE=Asia/Shanghai CODE_DIR=/code

RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

COPY ./ $CODE_DIR/

WORKDIR $CODE_DIR

CMD chmod +x entrypoint.sh && ./entrypoint.sh

EXPOSE 8080
