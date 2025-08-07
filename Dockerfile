FROM golang:1.24.3-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o ./tmp/main cmd/geohunter/main.go

FROM alpine:latest AS production

WORKDIR /app

COPY --from=build /app/tmp/main /app/main

CMD ["./main"]