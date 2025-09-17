FROM golang:1.25

WORKDIR /app

COPY go.mod ./
COPY . .

WORKDIR /app/cmd/comments
RUN go build -o comments .

CMD ["./comments"]
