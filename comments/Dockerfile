FROM golang:1.25

WORKDIR /app

COPY go.mod ./
COPY . .

WORKDIR /app/comments
RUN go mod tidy
RUN go build -o comments ./cmd/comments

CMD ["./comments"]