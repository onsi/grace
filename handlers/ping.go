package handlers

import "net/http"

type Ping struct {
}

func (p *Ping) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Grace abounds"))
	w.WriteHeader(http.StatusOK)
}
