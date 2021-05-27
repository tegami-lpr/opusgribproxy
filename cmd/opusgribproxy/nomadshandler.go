package main

import (
	"fmt"
	"net/http"
	"strings"
)

func fixDirPath(gfsPath string) string {
	return gfsPath[:len(gfsPath)-2] + "/" + gfsPath[len(gfsPath)-2:] + "/atmos"
}

func fixNomadsURL(gribUrl string) string {
	urlParts := strings.Split(gribUrl, "&")

	var gfsPath string
	var idx int

	if len(urlParts) == 35 {
		idx = 34
	} else {
		idx = 3
	}

	gfsPath = urlParts[idx]

	urlParts[idx] = fixDirPath(gfsPath)

	return strings.Join(urlParts, "&")
}

func nomadsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("Serve nomad GRIB for %s\n", r.URL.RawQuery)

	fixedUrlPart := fixNomadsURL(r.URL.RawQuery)
	fixedUrl := "https://nomads.ncep.noaa.gov/cgi-bin/filter_gfs_1p00.pl?" + fixedUrlPart

	fmt.Printf("Fixed URL is %s\n", fixedUrl)

	gribLen := getFetchURL(fixedUrl, w)
	if gribLen > -1 {
		fmt.Printf("Nomads GRIB for %s served. size %d\n", r.URL.RawQuery, gribLen)
	}
}
