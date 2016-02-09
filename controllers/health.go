package controllers

import (
	"fmt"
	"net/http"

	"github.com/zenazn/goji/web"
)

func (cg *ControllerGroup) Health(c web.C, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s\n", "Uh, we had a slight weapons malfunction, but uh... everything's perfectly all right now. We're fine. We're all fine here now, thank you. How are you?")
}
