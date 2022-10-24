FROM golang:alpine as build-stage

RUN mkdir /build
WORKDIR /build

COPY go.mod go.sum ./
RUN  go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go

FROM gcr.io/distroless/static-debian11 as prd

WORKDIR /root/

COPY --from=build-stage /build/main .

CMD ["./main"]
 