default: build

install_deps:
	go get -v ./...

build: install_deps
	go build -o ./bin/go-reverse-proxy .

clean:
	rm -rf ./bin
