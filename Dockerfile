FROM golang:latest as builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN make go-get
ENV GOBASE=/build
ENV GOPATH=$GOBASE/vendor:/$GOBASE
ENV GOBIN=$GOBASE/bin
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
COPY --from=builder /build/.env /app/
WORKDIR /app
CMD ["./main"]
