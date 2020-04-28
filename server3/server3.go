// server 3
// If you want to run it:
//   sudo lsof -i:8000 ##make sure nobody else is listening on port 8000
//   go run server3.go &
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"	
)

var mu sync.Mutex
var count int
var f = fmt.Fprintf

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handle echoes the http request
func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	f(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		f(w, "Header[%q] = %q\n", k, v)
	}
	f(w, "Host = %q\n", r.Host)
	f(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		f(w, "Form[%q] = %q\n", k, v)
	}
	
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request){
	mu.Lock()
	f(w, "Count %d\n", count)
	mu.Unlock()
}