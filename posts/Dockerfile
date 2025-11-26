FROM golang:1.25.4 AS build

WORKDIR /app

# Copiamos TODO el monorepo, no solo /users
COPY . .

# Entramos al microservicio específico
WORKDIR /app/posts

RUN go mod tidy
RUN go build -o posts ./cmd/posts

FROM gcr.io/distroless/base
COPY --from=build /app/posts/posts /posts
ENTRYPOINT ["/posts"]