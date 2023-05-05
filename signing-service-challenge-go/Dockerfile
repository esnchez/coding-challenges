FROM golang:1.16-alpine AS build
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/app .

EXPOSE 8080

CMD ["./app"]