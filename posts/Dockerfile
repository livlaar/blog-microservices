FROM golang:1.25

WORKDIR /app

COPY go.mod ./
COPY . .

WORKDIR /app/cmd/posts
RUN go build -o posts .

CMD ["./posts"]
