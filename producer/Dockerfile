# Dockerfile producer
FROM golang:1.18.4-alpine3.15 as build-stage

RUN mkdir /build
WORKDIR /build

COPY ./producer .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/main.go ./cmd/setup.go

FROM gcr.io/distroless/static-debian11 as prd

WORKDIR /root/
COPY --from=build-stage /build/main .
COPY --from=build-stage /build/.env .

EXPOSE 8080

CMD ["./main"]
 