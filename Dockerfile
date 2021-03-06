FROM golang:alpine AS builder
RUN apk update && apk add --no-cache git
ENV USER=appuser
ENV UID=100001
RUN adduser \
        --disabled-password \
        --gecos "" \
        --home "/whereever" \
        --shell "/sbin/nologin" \
        --no-create-home \
        --uid "${UID}" \
        "${USER}"
WORKDIR /app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /app/bin/thitcho

FROM scratch
WORKDIR /app
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
USER appuser:appuser
COPY --from=builder /app/bin/thitcho .
COPY --from=builder /app/config.yaml .
EXPOSE 6789

ENTRYPOINT ["/app/thitcho","server","--config","/app/config.yaml"]
