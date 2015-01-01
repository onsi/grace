package handlers

import (
	"net/http"
	"os"
)

type Cwd struct {
}

func (p *Cwd) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cwd, err := os.Getwd()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	out := "<dl>"
	out += "<dt>Current Working Directory</dt>"
	out += "<dd>" + cwd + "</dd>"
	out += "</dl>"
	styledTemplate.Execute(w, Body{`<div class="envs">` + out + `</div>`})
}
