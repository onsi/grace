package handlers

import (
	"net/http"
	"os"
	"strings"
)

type Env struct {
}

func (p *Env) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strings.Join(os.Environ(), "\n")))
	w.WriteHeader(http.StatusOK)
}
