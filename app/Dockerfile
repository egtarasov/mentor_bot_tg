FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN GOOS=linux go build -o bot ./cmd/main.go

COPY ./images ./config.yaml ./

CMD ["./bot"]