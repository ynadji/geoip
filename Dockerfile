FROM golang:1.10.4 AS builder

WORKDIR /go/src/github.com/ynadji/geoip
ADD . .

RUN go build -o /out/geoip

WORKDIR /out

RUN curl -k http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.mmdb.gz | gunzip - > GeoLite2-City.mmdb
EXPOSE 1234
ENTRYPOINT ["/out/geoip"]
