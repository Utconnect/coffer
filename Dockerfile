FROM golang:1.22.5-alpine

WORKDIR /app
EXPOSE 8080

COPY go.mod go.sum ./
RUN go mod download

COPY ./src .
RUN CGO_ENABLE=0 GOOS=linux go build -o ./coffer .

ENTRYPOINT ["./coffer"]