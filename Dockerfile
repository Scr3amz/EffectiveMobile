FROM golang:alpine as builder
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download -x

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /srv

COPY --from=builder /app/main .
COPY --from=builder /app/.env .

EXPOSE 8080

CMD sleep 5 && ./main