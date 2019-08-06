# VERSION 1.0.1
# AUTHOR: gum
# DESCRIPTION: dash server
# BUILD: docker build --rm -t registry.cn-shanghai.aliyuncs.com/iquant/dash-server:go1.0.1  .


FROM golang:latest

RUN mkdir -p $GOPATH/src/github.com
WORKDIR $GOPATH/src/LoopGraph
#ADD $GOPATH/src $GOPATH/src
COPY . $GOPATH/src/LoopGraph
RUN go-wrapper download
RUN go-wrapper install

#RUN go build .

EXPOSE 8070
ENTRYPOINT ["./LoopGraph"]