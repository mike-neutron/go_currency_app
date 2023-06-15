# Choose whatever you want, version >= 1.16
FROM golang:1.20.5-alpine

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

CMD ["air", "-c", ".air.toml"]
