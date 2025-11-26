FROM golang:1.25.4 AS build

WORKDIR /app

# Copiamos TODO el monorepo, no solo /users
COPY . .

# Entramos al microservicio específico
WORKDIR /app/comments

RUN go mod tidy
RUN go build -o comments ./cmd/comments

FROM gcr.io/distroless/base
COPY --from=build /app/comments/comments /comments
ENTRYPOINT ["/comments"]