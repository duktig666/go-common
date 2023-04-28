FROM golang:1.20-buster as builder
ENV GO111MODULE=on
ENV GOPATH=/go
ENV GOPROXY=https://goproxy.cn
RUN echo $GOPATH
WORKDIR /app
ADD . .
RUN CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build -o /go/bin/go-common  main.go

# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/base-debian10:debug-nonroot
WORKDIR /app
COPY --from=builder /go/bin/go-common .
ENTRYPOINT ["./go-common", "version"]
