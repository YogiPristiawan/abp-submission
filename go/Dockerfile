FROM golang:1.21.4-alpine3.18 AS build

WORKDIR /app

RUN apk add \
  build-base

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build \
  -tags timetzdata \
  -o binary \
  .

FROM alpine:3.18

RUN apk --no-cache add \
  ca-certificates \
  bash \
  curl

WORKDIR /app

COPY --from=build /app ./

EXPOSE 3030

ENTRYPOINT ["./binary"]
