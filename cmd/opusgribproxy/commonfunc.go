package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getFetchURL(url string, w http.ResponseWriter) int {
	ok := false
	gribLen := -1
	var errString string
	for {
		resp, err := http.Get(url)
		if err != nil {
			errString = err.Error()
			log.Fatal(err)
		}

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			errString = err.Error()
			log.Fatal(err)
		}
		gribLen = len(b)

		//resp.Header.Write(w)
		w.Write(b)
		ok = true
		break
	}

	if !ok {
		w.WriteHeader(http.StatusServiceUnavailable)
		fmt.Fprintf(w, "Error: %s", errString)
	}
	return gribLen
}
