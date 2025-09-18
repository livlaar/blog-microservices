FROM golang:1.25

WORKDIR /app

COPY go.mod ./
COPY . .

WORKDIR /app/users
RUN go mod tidy
RUN go build -o users ./cmd/users

CMD ["./users"]
