FROM golang:1.17 AS builder

WORKDIR /go/src/todolist-go-api
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app ./cmd/todolist

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/todolist-go-api/app .
CMD ["./app"]
