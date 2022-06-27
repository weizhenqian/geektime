FROM golang:1.18-alpine AS build
ENV GOPATH=''
ENV CGO_ENABLED=0
ENV GOPROXY=https://goproxy.cn
COPY main.go /go/src
WORKDIR /go/src
RUN go mod init homework6 && \
  go mod tidy && \
  go build -o /bin/homework6

FROM scratch
COPY --from=build /bin/homework6 /bin/homework6
ENTRYPOINT ["/bin/homework6"]
