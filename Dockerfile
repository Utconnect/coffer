FROM golang:1.22.5-alpine
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./src .
RUN CGO_ENABLE=0 GOOS=linux go build -o ./coffer .

ENV SERVICE_NAME="coffer"
RUN addgroup --gid 1001 -S "$SERVICE_NAME" && \
    adduser -G "$SERVICE_NAME" --shell /bin/false --disabled-password -H --uid 1001 "$SERVICE_NAME" && \
    mkdir -p /var/log/"$SERVICE_NAME" && \
    chown "$SERVICE_NAME":"$SERVICE_NAME" /var/log/"$SERVICE_NAME"

EXPOSE 8080
USER "$SERVICE_NAME"

ENTRYPOINT ["./coffer"]