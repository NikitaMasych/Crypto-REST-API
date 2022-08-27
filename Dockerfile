FROM golang:1.18.4-alpine3.15 as build-stage

RUN mkdir /app
WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main

FROM golang:1.18.4-alpine3.15 as prd
COPY --from=build-stage /app/main /app/main

EXPOSE 8080

CMD ["/app/main"]
