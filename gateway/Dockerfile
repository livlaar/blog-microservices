FROM golang:1.25.4 AS build

WORKDIR /app

# Copia gateway y shared solamente
COPY gateway ./gateway
COPY shared ./shared

WORKDIR /app/gateway

RUN go mod tidy
RUN go build -o gateway .

FROM gcr.io/distroless/base
COPY --from=build /app/gateway/gateway /gateway
ENTRYPOINT ["/gateway"]
