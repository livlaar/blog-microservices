FROM golang:1.25.4 AS build

WORKDIR /app

# Copiamos TODO el monorepo, no solo /users
COPY . .

# Entramos al microservicio específico
WORKDIR /app/users

RUN go mod tidy
RUN go build -o users ./cmd/users

FROM gcr.io/distroless/base
COPY --from=build /app/users/users /users
ENTRYPOINT ["/users"]
