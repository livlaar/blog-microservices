FROM golang:1.25

WORKDIR /app

COPY go.mod ./
COPY . .

WORKDIR /app/cmd/users
RUN go build -o users .

CMD ["./users"]