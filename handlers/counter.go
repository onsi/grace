package handlers

import (
	"fmt"
	"net/http"
	"sync/atomic"
)

type IncrementCounter struct {
	SharedCounter *uint64
}

func (p *IncrementCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	atomic.AddUint64(p.SharedCounter, 1)
}

type ReadCounter struct {
	SharedCounter *uint64
}

func (p *ReadCounter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%d", atomic.LoadUint64(p.SharedCounter))
}
