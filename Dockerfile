FROM golang:1.8

WORKDIR /go/src/github.com/nowenl/kecd

COPY . .

RUN go build -o kecd github.com/nowenl/kecd/cmd

EXPOSE 8876
CMD ["./kecd", "server"]