all: clean release
build:
	GOOS=linux GOARCH=amd64 go build -o main main.go
release: build
	zip main.zip main
clean:
	rm -f main.zip main
