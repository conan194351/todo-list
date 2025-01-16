FROM golang:1.23.4-alpine AS builder

WORKDIR /app
ENV GOPATH /go
ENV CGO_ENABLED 0
ENV GO111MODULE on
ENV GOPROXY https://proxy.golang.org

RUN apk add gcc && \
    apk add make

COPY go.mod go.sum  ./

RUN go mod download && go mod verify && go mod tidy

COPY . ./

RUN go build -o /app/todolist main.go

FROM golang:1.23.4-alpine

WORKDIR /root

COPY --from=builder /app/todolist .

EXPOSE 8080

CMD ["./todolist","api:launch"]