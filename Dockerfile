# ref: https://docs.docker.com/language/golang/build-images/

# Build
FROM golang:1.17-bullseye AS build

WORKDIR /src

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /server ./cmd/todolist

# Deploy
FROM gcr.io/distroless/base-debian11

COPY --from=build /server /server
CMD ["/server"]
