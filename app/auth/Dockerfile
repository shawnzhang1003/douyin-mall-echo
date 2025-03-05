FROM golang:1.23-alpine AS build

WORKDIR /app
RUN export PORT=8080 APP_ENV=local DB_HOST=localhost DB_PORT=3306 DB_DATABASE=users DB_USERNAME=dbuser DB_PASSWORD=password1234 DB_ROOT_PASSWORD=password4321

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go

FROM alpine:3.20.1 AS prod
WORKDIR /app
COPY --from=build /app/main /app/main
EXPOSE ${PORT}
CMD ["./main"]


