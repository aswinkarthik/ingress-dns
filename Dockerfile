FROM golang:1.9.2 as source
RUN mkdir -p /go/src/github.com/aswinkarthik93
COPY . /go/src/github.com/aswinkarthik93/ingress-dns
RUN cd /go/src/github.com/aswinkarthik93/ingress-dns \
    && go build -o ingress-dns

FROM debian:stretch
RUN apt-get update -y \
    && apt-get install ca-certificates -y
RUN update-ca-certificates --verbose
COPY --from=source /go/src/github.com/aswinkarthik93/ingress-dns/ingress-dns .
CMD ["./ingress-dns", "start"]
