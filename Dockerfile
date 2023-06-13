FROM golang:1.20-alpine as builder

RUN apk add --no-cache make

WORKDIR /app/go-reverse-proxy

COPY . .

RUN make build

# ---

FROM alpine:3.15

COPY --from=builder /app/go-reverse-proxy/bin/go-reverse-proxy /usr/bin/go-reverse-proxy

CMD sh -c "go-reverse-proxy --origin=$ORIGIN --host=$HOST --port=$PORT"
