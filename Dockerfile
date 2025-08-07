FROM golang:1.24.3-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main server.go

EXPOSE 3000

CMD ["./main"]

FROM openresty/openresty:alpine-fat

RUN apk add --no-cache git build-base && \
    luarocks install lua-cjson && \
    luarocks install lua-resty-jwt


