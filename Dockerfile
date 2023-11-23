# stage 1: compile binary
FROM golang:1.21.1-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /app/chatting_app ./cmd/main.go

EXPOSE 8080

CMD ["./chatting_app"]