FROM golang:latest as builder

ENV GO111MODULE=auto
ENV SERVER_HOST=grpc-server:50051

WORKDIR /go/src/github.com/maei/golang_grpc_sum_client

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o grpc-client /go/src/github.com/maei/golang_grpc_sum_client/src/main.go
EXPOSE 8002

# CMD ["./grpc-client"]

FROM alpine
COPY --from=builder /go/src/github.com/maei/golang_grpc_sum_client/grpc-client /
CMD ["./grpc-client"]