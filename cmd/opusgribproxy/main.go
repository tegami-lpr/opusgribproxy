package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusServiceUnavailable)
	fmt.Fprintf(w, "%d: No service for %s", http.StatusServiceUnavailable, r.URL.Path[1:])
	fmt.Printf("%d: No service for %s\n", http.StatusServiceUnavailable, r.URL.Path)
}

func main() {
	fmt.Println("Opus proxy started...")

	//http.HandleFunc("/data/nccf/com/gfs/prod/", gribRedirectHandler)
	http.HandleFunc("/data/nccf/com/gfs/prod/", gribHandler)
	http.HandleFunc("/cgi-bin/filter_gfs_1p00.pl", nomadsHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
