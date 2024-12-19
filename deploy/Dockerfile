FROM golang:1.23.3-alpine3.19 AS builder

RUN apk --no-cache update && apk add --no-cache build-base git

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

ARG BUILD_VERSION=unknown

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=${BUILD_VERSION}'" -a -installsuffix cgo -o main .

FROM alpine:3.20

WORKDIR /app

RUN apk add --no-cache --upgrade tzdata

COPY --from=builder /app/main .

COPY --from=builder /app/internal/domain/migrations /app/internal/domain/migrations

ENTRYPOINT ["./main"]