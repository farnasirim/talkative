package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var startTime int64 = time.Now().UnixNano()

const (
	Nanosecond  int64 = 1
	Microsecond       = 1000 * Nanosecond
	Milisecond        = 1000 * Microsecond
	Second            = 1000 * Milisecond
	Minute            = 60 * Second
	Hour              = 60 * Minute
	Day               = 24 * Hour
	Year              = 365 * Day
)

func humanReadable(tm int64) string {
	var build []string
	durationsValues := []int64{Milisecond, Second, Minute, Hour, Day, Year}
	durationsNames := []string{"milisecond", "second", "minute", "hour", "day", "year"}

	for i := len(durationsValues) - 1; i >= 0; i-- {
		numOfThisUnit := tm / durationsValues[i]
		tm %= durationsValues[i]
		if numOfThisUnit != 0 {
			suffix := "s"
			if numOfThisUnit == 1 {
				suffix = ""
			}
			build = append(build, fmt.Sprintf("%d %s%s", numOfThisUnit, durationsNames[i], suffix))
		}
	}

	return strings.Join(build, " ")
}

func info(w http.ResponseWriter, r *http.Request) {
	var toWrite bytes.Buffer
	hostname, err := os.Hostname()

	if err != nil {
		toWrite.Write([]byte("Could not read hostname"))
	} else {
		toWrite.Write([]byte("This is "))
		toWrite.Write([]byte(hostname))
	}
	toWrite.Write([]byte("\n"))

	toWrite.Write([]byte("Host: "))
	toWrite.Write([]byte(r.Host))
	toWrite.Write([]byte("\n"))

	toWrite.Write([]byte("Requested path: "))
	toWrite.Write([]byte(r.URL.EscapedPath()))
	toWrite.Write([]byte("\n"))

	toWrite.Write([]byte("Remote Address: "))
	toWrite.Write([]byte(r.RemoteAddr))
	toWrite.Write([]byte("\n"))

	toWrite.Write([]byte("Uptime: "))
	toWrite.Write([]byte(humanReadable(time.Now().UnixNano() - startTime)))
	toWrite.Write([]byte("\n"))

	toWrite.Write([]byte("\n"))

	w.Write(toWrite.Bytes())
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", info)
	log.Fatalln(http.ListenAndServe(":80", mux))
}
