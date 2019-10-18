FROM golang:1.13

#run mkdir /go/src/github.com
#run mkdir /go/src/github.com/RESTApp

WORKDIR /go/src/RESTApp
COPY . .

RUN go get -v ./...
RUN go build .
EXPOSE 8081

CMD ["/main"]
