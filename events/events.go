package main

import (
  "fmt"
  "net/http"
)

type Event struct {
  name string
}

func init() {
  eventList = make([]Event, 0)
}

var eventList []Event

func putEvent(e Event) {
  eventList = append(eventList , e) 
}

func events(w http.ResponseWriter, r *http.Request) {
  if r.Method == "PUT" {
    event := Event{r.FormValue("name")}
    putEvent(event)
  } else {
    fmt.Fprintf(w, "%v", eventList)
  }
}

func main() {
  http.HandleFunc("/events", events)
  http.ListenAndServe(":8080", nil)
}
