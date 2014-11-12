package handlers

import (
	"fmt"
	"net/http"
	"time"
)

type StartedAt struct {
	Time time.Time
}

func (p *StartedAt) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%d", p.Time.UnixNano())))
}
