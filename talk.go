package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/colonelmo/talkative/web"
)

func handleSignals() {
	signalsChannel := make(chan os.Signal, 1)
	signal.Notify(signalsChannel, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	for signal := range signalsChannel {
		fmt.Println("received signal:")
		fmt.Println(signal.String())
		fmt.Println()
	}
}

func main() {
	go handleSignals()

	server := web.Server{StartTime: time.Now().UnixNano()}

	mux := http.NewServeMux()
	mux.HandleFunc("/", server.Info)
	mux.HandleFunc("/count", server.Count)
	mux.HandleFunc("/count/reset", server.CountReset)

	log.Fatalln(http.ListenAndServe(":80", mux))
}
