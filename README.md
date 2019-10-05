# geoip

`geoip` provides IP geolocation as a service, e.g., converting IPv4 dotted-quad
addresses (`a.b.c.d`) into a latitude and longitude. It is composed of two
containerized services and orchestrated through `docker-compose`. The two services are:

* `geoip`: a service in Go that translates IPs to a geographic location, and
* `webapp`: a PHP application running on Apache as a front-end.

## Building & Running

```
$ git clone git@github.com:ynadji/geoip.git
$ cd geoip
$ docker-compose build
$ docker-compose up
```

This will start the two aforementioned services on the IP address used by docker
on port 80. For my `docker-machine` install, the IP is `192.168.99.100`. Below
shows an interaction with the services at the command-line, however, navigating
to the IP in your browser will provide an even uglier UI.

```
$ curl "http://192.168.99.100?ip=$(dig +short sif.nadji.us)"

<html>
  <head>
    <title>GIP GIP</title>
  </head>
  <body>
    <form action="index.php" method="get">
      <input id="ip" type="text" name="ip" value="1.2.3.4"></input>
      <input id="submit" type="submit" value="FINDIT" />
    </form>
    <div id="geo">Location for 138.88.4.99 is: {"latitude":38.9208,"longitude":-77.036}</div>
  </body>
</html>
```

## Testing

Tests are automatically run at build time using Go's built-in testing
functionality. See `geoip_test.go` for details. Tests can be run independently
with:

```
$ go get github.com/ynadji/geoip
$ go test github.com/ynadji/geoip
```

Note the MaxMind DB must be downloaded, decompressed, and in the current working
directory for the tests to succeed as DB initialization is tested.
