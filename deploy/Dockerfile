FROM golang:1.23.3-alpine3.19 as builder

RUN set -ex \
    && apk --no-cache update \
    && apk add --no-cache git curl openssh-client ca-certificates \
       make perl bash build-base zlib-dev ucl-dev \
    && rm -rf /var/cache/apk/*

WORKDIR /go/src/ohmytp

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

ARG BUILD_VERSION=unknown

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-X 'main.Version=${BUILD_VERSION}'" -a -installsuffix cgo -o main .


FROM alpine:3.20

WORKDIR /go/src/ohmytp

RUN apk add --no-cache --upgrade \
    tzdata

COPY --from=builder /go/src/ohmytp/main .

COPY --from=builder /go/src/ohmytp/internal/domain/migrations /go/src/ohmytp/internal/domain/migrations

ENTRYPOINT ["./main"]