package handlers

import (
	"net/http"
	"os"
)

type InstanceIndex struct {
}

func (p *InstanceIndex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(os.Getenv("INSTANCE_INDEX")))
}
