package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/zenazn/goji/web"
)

// Database handles requests for a dump of the loot database
func (cg *ControllerGroup) DumpDatabase(c web.C, w http.ResponseWriter, r *http.Request) {
	latitude, err := strconv.ParseFloat(c.URLParams["latitude"], 64)
	if err != nil {
		log.Panic(err)
	}

	longitude, err := strconv.ParseFloat(c.URLParams["longitude"], 64)
	if err != nil {
		log.Panic(err)
	}

	radius, err := strconv.ParseFloat(c.URLParams["radius"], 64)
	if err != nil {
		log.Panic(err)
	}

	data, err := cg.Accessors.DumpDatabase(latitude, longitude, radius)
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprintf(w, "%s\n", data)
}
