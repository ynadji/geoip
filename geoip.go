package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/oschwald/geoip2-golang"
)

var db *geoip2.Reader
var logger = log.New(os.Stderr, "", 0)

type geo struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func initdb() {
	var err error
	db, err = geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		db = nil
		logger.Panic("Failed to open MaxMind DB: " + err.Error())
	}
}

func geoFromIP(ip net.IP) (geo, error) {
	if db == nil {
		logger.Panic("DB not initialized")
	}

	r, err := db.City(ip)
	if err != nil {
		logger.Printf("Failed to lookup IP: %+v", err)
		return geo{}, err
	}

	return geo{r.Location.Latitude, r.Location.Longitude}, nil
}

func geoFromIPString(ip string) (geo, error) {
	ipaddr := net.ParseIP(ip)
	if ipaddr == nil {
		errmsg := fmt.Sprintf("Failed to parse IP: %+v", ip)
		logger.Printf(errmsg)
		return geo{}, errors.New(errmsg)
	}
	return geoFromIP(ipaddr)
}

func main() {
	if len(os.Args) != 1 {
		logger.Panic("usage: go run geoip")
	}
	initdb()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		ips, ok := r.URL.Query()["ip"]

		if !ok || len(ips[0]) < 1 {
			errmsg := "URL parameter 'ip' is missing"
			logger.Print(errmsg)
			fmt.Fprint(w, errmsg)
			return
		}
		geo, err := geoFromIPString(ips[0])
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		byteResp, err := json.Marshal(geo)
		if err != nil {
			panic("Failed to marshall tree")
		}
		fmt.Fprintf(w, string(byteResp))
	})

	fmt.Println("Starting service...")
	http.ListenAndServe(":1234", nil)
}
