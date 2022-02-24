ARG GO_VERSION=1.17

FROM golang:${GO_VERSION} AS builder
LABEL MAINTAINER=ugurluerdi@gmail.com
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group
WORKDIR /go/src/helloworld

ENV GO111MODULE=on
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY main.go .
COPY VERSION .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


FROM alpine:latest AS final
# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/
WORKDIR /opt/
COPY --from=builder /go/src/helloworld/app .
COPY --from=builder /go/src/helloworld/VERSION .
USER nobody:nobody
CMD ["./app"]
EXPOSE 8080