# Dockerfile consumer
FROM golang:1.18.4-alpine3.15 as build-stage

RUN mkdir /build
WORKDIR /build

COPY ./consumer .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM gcr.io/distroless/static-debian11 as prd

WORKDIR /root/
COPY --from=build-stage /build/main .

CMD ["./main"]
 