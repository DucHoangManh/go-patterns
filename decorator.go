package gopatterns

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func runServer() {
	http.HandleFunc("/", durationLogger(helloEndpoint))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world")
}

func durationLogger(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		startTime := time.Now()
		f(writer, request)
		log.Printf("complete handle request after %d us", time.Since(startTime).Microseconds())
	}
}
