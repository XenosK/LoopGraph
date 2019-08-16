# VERSION 1.0.1
# AUTHOR: gum
# DESCRIPTION: dash server
<<<<<<< HEAD
# BUILD: docker build --rm -t registry.cn-shanghai.aliyuncs.com/iquant/dash-server:go1.0.1  .


FROM golang:latest

WORKDIR $GOPATH/src/LoopGraph

COPY . $GOPATH/src/LoopGraph
RUN go get github.com/GumKey/LoopGraph

#RUN go build .
=======
# BUILD: docker build --rm -t registry.cn-shanghai.aliyuncs.com/iquant/dash-server:go1.2.1 .


#FROM golang:alpine
FROM gowebtest:latest

WORKDIR /opt/LoopGraph/
COPY . /opt/LoopGraph/
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn
#RUN apk add --no-cache gcc musl-dev git && go build -i -v
RUN go build -i -v

>>>>>>> update laui html and json

EXPOSE 8070
ENTRYPOINT ["./LoopGraph"]