package controllers

import (
	"fmt"
	"net/http"

	"github.com/jessemillar/rytsar-server/helpers"
	"github.com/zenazn/goji/web"
)

// Collect handles http requests to collect loot at a given location
func (cg *ControllerGroup) Collect(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", helpers.Collect(r.PostFormValue("username"), r.PostFormValue("latitude"), r.PostFormValue("longitude"), cg.Accessors))
}
