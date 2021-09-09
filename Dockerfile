FROM golang:1.17 AS builder

RUN apt-get install -y git

WORKDIR /opt/src/
ADD . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -o goas .

FROM alpine:3.9

RUN apk --no-cache add ca-certificates

WORKDIR /opt

COPY --from=builder /opt/src/goas /bin

RUN chmod +x /bin/goas

ENTRYPOINT "goas"