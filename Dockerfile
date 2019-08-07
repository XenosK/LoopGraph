# VERSION 1.0.1
# AUTHOR: gum
# DESCRIPTION: dash server
# BUILD: docker build --rm -t registry.cn-shanghai.aliyuncs.com/iquant/dash-server:go1.0.1  .


FROM golang:latest

WORKDIR $GOPATH/src/LoopGraph

COPY . $GOPATH/src/LoopGraph
RUN go get github.com/GumKey/LoopGraph

#RUN go build .

EXPOSE 8070
ENTRYPOINT ["./LoopGraph"]