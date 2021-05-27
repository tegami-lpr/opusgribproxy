package main

import (
	"fmt"
	"net/http"
	"strings"
)

func fixGFSPath(gfsPath string) string {
	return gfsPath[:len(gfsPath)-2] + "/" + gfsPath[len(gfsPath)-2:] + "/atmos"
}

func fixGribURL(gribUrl string) string {
	urlParts := strings.Split(gribUrl, "/")

	gfsPath := urlParts[5]
	urlParts[5] = fixGFSPath(gfsPath)

	return strings.Join(urlParts, "/")
}

func gribHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serve GRIB for %s\n", r.URL.Path[1:])

	fixedUrl := "http://ftp.ncep.noaa.gov/" + fixGribURL(r.URL.Path[1:])
	fmt.Printf("Fixed URL is %s\n", fixedUrl)

	gribLen := getFetchURL(fixedUrl, w)

	if gribLen > -1 {
		fmt.Printf("GRIB for %s served. size %d\n", r.URL.Path[1:], gribLen)
	}
}

func gribRedirectHandler(w http.ResponseWriter, r *http.Request) {
	fixedUrl := "http://ftp.ncep.noaa.gov/" + fixGribURL(r.URL.Path[1:])

	fmt.Printf("Serve GRIB for %s\n", r.URL.Path[1:])
	fmt.Printf("Redirect to %s\n", fixedUrl)

	w.Header().Set("Location", fixedUrl)
	w.WriteHeader(http.StatusPermanentRedirect)
}
