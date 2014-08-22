package handlers

import (
	"net/http"
	"os"
	"strings"
)

type Env struct {
}

func (p *Env) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	out := "<dl>"
	for _, env := range os.Environ() {
		kv := strings.Split(env, "=")
		out += "<dt>" + kv[0] + "</dt>"
		out += "<dd>" + kv[1] + "</dd>"
	}
	out += "</dl>"
	styledTemplate.Execute(w, Body{out})
}
