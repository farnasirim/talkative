package web

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-yaml/yaml"

	"github.com/colonelmo/talkative/utils"
)

type Server struct {
	StartTime int64
	counter   int
	total     int
}

func (s *Server) Info(w http.ResponseWriter, r *http.Request) {
	s.total++
	respo := InfoResponse{}
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Could not read hostname"
	}
	respo.Hostname = hostname

	respo.RequestedPath = r.URL.EscapedPath()
	respo.RemoteAddress = r.RemoteAddr
	now := time.Now().UnixNano()
	respo.TotalNumberOfRequests = s.total
	respo.CurrentTimeUnixNano = now
	respo.UptimeHumanReadable = utils.HumanReadable(now - s.StartTime)
	respo.UptimeUnixNano = now - s.StartTime

	marshalledResponse, err := yaml.Marshal(&respo)
	fmt.Println(string(marshalledResponse))
	w.Write([]byte("v1.1\n"))
	w.Write(marshalledResponse)
}

func (s *Server) Count(w http.ResponseWriter, r *http.Request) {
	s.total++
	s.counter++
	respo := CountResponse{}
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Could not read hostname"
	}
	respo.Hostname = hostname

	respo.RequestedPath = r.URL.EscapedPath()
	respo.TotalNumberOfRequests = s.total
	respo.Count = s.counter
	now := time.Now().UnixNano()
	respo.CurrentTimeUnixNano = now
	respo.UptimeHumanReadable = utils.HumanReadable(now - s.StartTime)
	respo.UptimeUnixNano = now - s.StartTime

	marshalledResponse, err := yaml.Marshal(&respo)
	fmt.Println(string(marshalledResponse))
	w.Write(marshalledResponse)
}

func (s *Server) CountReset(w http.ResponseWriter, r *http.Request) {
	s.total++
	respo := CountResetResponse{}
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Could not read hostname"
	}
	respo.Hostname = hostname

	respo.RequestedPath = r.URL.EscapedPath()
	respo.Count = s.counter

	setCounterTo, err := strconv.ParseInt(r.FormValue("value"), 10, 64)
	if err != nil {
		setCounterTo = 0
	}
	s.counter = int(setCounterTo)

	respo.TotalNumberOfRequests = s.total
	respo.SetTo = s.counter
	now := time.Now().UnixNano()
	respo.CurrentTimeUnixNano = now
	respo.UptimeHumanReadable = utils.HumanReadable(now - s.StartTime)
	respo.UptimeUnixNano = now - s.StartTime

	marshalledResponse, err := yaml.Marshal(&respo)
	fmt.Println(string(marshalledResponse))
	w.Write(marshalledResponse)
}
