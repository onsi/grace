package handlers

import (
	"net/http"
	"os"
)

type InstanceIndex struct {
}

func (p *InstanceIndex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	index := os.Getenv("CF_INSTANCE_INDEX")
	w.Write([]byte(index))
}
