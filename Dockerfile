FROM golang:1.13.0 as builder
WORKDIR /go/src/github.com/BaronMsk/ssl-checker
COPY ./ /go/src/github.com/BaronMsk/ssl-checker
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:3.10
COPY --from=builder /go/src/github.com/BaronMsk/ssl-checker /ssl-checker
ENTRYPOINT ["ssl-checker"]