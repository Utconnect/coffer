FROM golang:1.22.5
USER $APP_UID
WORKDIR /app
EXPOSE 8080

COPY go.mod ./
RUN go mod download

COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o /coffer
RUN chmod +x /coffer

CMD ["/coffer"]
