BINARY_NAME=RESTApp

hello:
	echo: "Hello"

build:
	go build 

clean:
	rm -f $(BINARY_NAME)

run:
	go run main.go

compile:
	echo "Compiling for Every OS & Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go