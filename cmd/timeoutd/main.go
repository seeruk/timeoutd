package main

import (
	"flag"
	"log"
	"net/http"
	"time"
)

const (
	// yearsSinceBirthOfUniverse is the number of years since the big bang, give or take a few.
	yearsSinceBirthOfUniverse = 14e9
	// yearToHourRatio is the number of hours in a year, roughly.
	yearToHourRatio = 8760
)

func main() {
	var addr string

	flag.StringVar(&addr, "addr", ":8080", "Address to listen on")
	flag.Parse()

	err := http.ListenAndServe(addr, http.HandlerFunc(handler))
	if err != nil {
		log.Fatal(err)
	}
}

// handler is an http.HandlerFunc that will sleep for a little while.
func handler(_ http.ResponseWriter, _ *http.Request) {
	for i := 0; i < yearsSinceBirthOfUniverse*yearToHourRatio; i++ {
		time.Sleep(time.Hour)
	}
}
