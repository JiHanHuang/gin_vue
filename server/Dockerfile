FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/github.com/JiHanHuang/gin_vue
COPY . $GOPATH/src/github.com/JiHanHuang/gin_vue
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./go-gin-example"]
