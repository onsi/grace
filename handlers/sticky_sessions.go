package handlers

import (
	"fmt"
	"net/http"

	"github.com/onsi/grace/helpers"
)

type Stick struct {
}

func (p *Stick) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name: "JSESSIONID",
	})

	index, _ := helpers.FetchIndex()
	w.Write([]byte(fmt.Sprintf("Stuck to %d", index)))
}

type Unstick struct {
}

func (p *Unstick) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   "JSESSIONID",
		MaxAge: -1,
	})
	w.Write([]byte(fmt.Sprintf("Unstuck")))
}
