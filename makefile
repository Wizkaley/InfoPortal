BINARY_NAME=RESTApp

hello:
	echo: "Hello"

build:
	go build 

clean:
	rm -f $(BINARY_NAME)

run:
	go run main.go
