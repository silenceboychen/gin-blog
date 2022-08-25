FROM golang:latest

WORKDIR $GOPATH/src/gin-blog
COPY . $GOPATH/src/gin-blog
RUN go build .

EXPOSE 8080
ENTRYPOINT ["./gin-blog"]