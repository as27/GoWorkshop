package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

var port = ":8080"

func main() {
	fmt.Println("Starte Server auf Port", port)
	http.HandleFunc("/test", testHandler)
	timer := &TimeHandler{}
	http.Handle("/time", timer)
	log.Fatal(http.ListenAndServe(port, nil))
}

type TimeHandler struct {
	RequestNr int       `json:"request_nr,omitempty"`
	Time      time.Time `json:"time,omitempty"`
}

func (th *TimeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	th.RequestNr++
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	th.Time = time.Now()
	e.Encode(th)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	b := bytes.NewBufferString("Test")
	b.WriteTo(w)
}
