FROM golang:1.20.5-alpine3.18
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy
COPY . .
RUN go build -o authapi ./cmd
EXPOSE 8080
CMD ["./authapi"]

