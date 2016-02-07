package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/zenazn/goji/web"
)

// Database handles requests for a dump of the loot database
func (cg *ControllerGroup) DumpDatabase(c web.C, w http.ResponseWriter, r *http.Request) {
	data, err := cg.Accessors.DumpDatabase()
	if err != nil {
		log.Panic(err)
	}

	fmt.Fprintf(w, "%s\n", data)
}
