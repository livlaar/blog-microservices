FROM golang:1.25

WORKDIR /app

COPY go.mod ./
COPY . .

WORKDIR /app/posts
RUN go mod tidy
RUN go build -o posts ./cmd/posts

CMD ["./posts"]
