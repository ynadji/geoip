package main

import (
	"net"
	"testing"
)

func TestInit(t *testing.T) {
	// Make sure the MaxMind DB can be initialized
	initdb()
}

func TestGeoFromIP(t *testing.T) {
	type testCase struct {
		ip       net.IP
		expected geo
		errIsNil bool
	}

	testCases := []testCase{
		{net.ParseIP("0.0.0.0"), geo{0.0, 0.0}, true},
		{net.ParseIP("1.2.3.4"), geo{37.751, -97.822}, true},
		{net.ParseIP("255.255.255.255"), geo{0.0, 0.0}, true},
	}

	for _, tc := range testCases {
		geo, err := geoFromIP(tc.ip)
		if ((tc.errIsNil && err != nil) || (!tc.errIsNil && err == nil)) || geo != tc.expected {
			t.Fatalf("Test case %+v failed, got %+v and %+v", tc, geo, err)
		}
	}
}

func TestGeoFromIPString(t *testing.T) {
	type testCase struct {
		ip       string
		expected geo
		errIsNil bool
	}

	testCases := []testCase{
		{"0.0.0.0", geo{0.0, 0.0}, true},
		{"1.2.3.4", geo{37.751, -97.822}, true},
		{"255.255.255.255", geo{0.0, 0.0}, true},
		{"notanip", geo{0.0, 0.0}, false},
	}

	for _, tc := range testCases {
		geo, err := geoFromIPString(tc.ip)
		if ((tc.errIsNil && err != nil) || (!tc.errIsNil && err == nil)) || geo != tc.expected {
			t.Fatalf("Test case %+v failed, got %+v and %+v", tc, geo, err)
		}
	}
}
