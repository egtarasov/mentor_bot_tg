FROM golang:1.22

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./
RUN GOOS=linux go build -o calendar ./main.go
CMD ["./calendar"]