package controllers

import (
	"fmt"
	"net/http"

	"github.com/jessemillar/rytsar-server/helpers"
	"github.com/zenazn/goji/web"
)

// Database handles requests for a dump of the loot database
func (cg *ControllerGroup) DumpDatabase(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", helpers.DumpDatabase(cg.Accessors))
}
