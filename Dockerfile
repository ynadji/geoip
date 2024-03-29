FROM golang:1.13.1 AS builder

WORKDIR /go/src/github.com/ynadji/geoip
ADD . .

RUN go get github.com/oschwald/geoip2-golang
RUN go build -o /out/geoip

WORKDIR /out

RUN curl -k http://geolite.maxmind.com/download/geoip/database/GeoLite2-City.mmdb.gz | gunzip - > GeoLite2-City.mmdb
# Link so tests will work. There's probably a better way to do this.
RUN ln GeoLite2-City.mmdb /go/src/github.com/ynadji/geoip/GeoLite2-City.mmdb
RUN go test github.com/ynadji/geoip
EXPOSE 1234
ENTRYPOINT ["/out/geoip"]
