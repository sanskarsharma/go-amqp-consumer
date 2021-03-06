FROM golang:1.14-alpine as builder
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY . .

ENV GOOS=linux \
    GOARCH=amd64 \
    GOBIN=$GOPATH/bin
RUN go mod download
RUN go build -ldflags="-s -w" -o ./bin/main-bin ./main.go

# creating an alpine image from scratch (lightweight)
FROM alpine:3.9

# copying binary built from previous stage
WORKDIR /usr/bin
COPY --from=builder /go/src/app/bin /go/bin
ENTRYPOINT /go/bin/main-bin 