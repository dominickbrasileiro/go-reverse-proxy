# go-reverse-proxy
"go-reverse-proxy" is a command line tool that creates a reverse proxy between your host and a origin.
You can use it to forward incoming requests to a server in another URL, for example.

## Usage
```
Usage:
  go-reverse-proxy [flags]

Flags:
  -h, --help            help for go-reverse-proxy
  -H, --host string     Host to listen to (default "0.0.0.0")
  -o, --origin string   Origin server URL
  -p, --port string     Port to listen to (default "8080")
```

### Example
Redirects http requests from `http://127.0.0.1:5000` -> `https://google.com`.
```sh
$ go-reverse-proxy --origin=https://google.com --host=127.0.0.1 --port=5000
```

## Installation
### Docker and Docker Compose
The "go-reverse-proxy" docker image is available upon [ghcr.io](https://github.com/dominickbrasileiro/go-reverse-proxy/pkgs/container/go-reverse-proxy) registry.
You can pull it running the command below:
```sh
$ docker pull ghcr.io/dominickbrasileiro/go-reverse-proxy
```

If you want to pull a specific version of "go-reverse-proxy", use a tag:
```sh
$ docker pull ghcr.io/dominickbrasileiro/go-reverse-proxy:1.1.0
```

To run the "go-reverse-proxy" image, you'll need to provide the following environment variables:
- HOST
- ORIGIN
- PORT

Example using `docker run` command:
```sh
$ docker run \
-e HOST=127.0.0.1 \
-e ORIGIN=https://google.com \
-e PORT=5000 \
--name go-reverse-proxy ghcr.io/dominickbrasileiro/go-reverse-proxy
```

Example using `docker-compose.yml`:
```yml
version: "3.8"
services:
  go-reverse-proxy:
    container_name: "go-reverse-proxy"
    image: "ghcr.io/dominickbrasileiro/go-reverse-proxy"
    restart: "always"
    environment:
      - "HOST=127.0.0.1"
      - "ORIGIN=https://google.com"
      - "PORT=5000"
```

### Building from source
#### Requirements
- [Go](https://go.dev/) (1.20 or later)

1. Using make and Makefile
```sh
$ git clone https://github.com/dominickbrasileiro/go-reverse-proxy.git
$ cd go-reverse-proxy
$ make build
```

2. Without Makefile
```sh
$ git clone https://github.com/dominickbrasileiro/go-reverse-proxy.git
$ cd go-reverse-proxy
$ go get -v ./...
$ go build -o ./bin/go-reverse-proxy .
```

Now you have the "go-reverse-proxy" binary in `./bin/go-reverse-proxy`. If you want to use it globally, run the command below:
```sh
$ sudo mv ./bin/go-reverse-proxy /usr/local/bin
```

---

Made with ❤️ by Dominick Brasileiro.

Feel free [to reach out](https://www.linkedin.com/in/dominickbrasileiro/)!

[![Linkedin Badge](https://img.shields.io/badge/-LinkedIn-blue?style=flat-square&logo=Linkedin&logoColor=white&link=https://www.linkedin.com/in/dominickbrasileiro/)](https://www.linkedin.com/in/dominickbrasileiro/)
