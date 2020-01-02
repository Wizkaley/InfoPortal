FROM golang:1.13

WORKDIR /go/src/InfoPortal
COPY . .

RUN go get -v ./...
RUN go build .
#RUN go test ./... -v
EXPOSE 8081

CMD ["./RESTApp"]
