FROM golang:1.23.0

WORKDIR /app

COPY go.mod go.sum ./

ENV GOPROXY=https://goproxy.cn,direct

RUN go mod download

COPY . .

RUN go build -o main ./cmd/api

CMD ["./main"]
