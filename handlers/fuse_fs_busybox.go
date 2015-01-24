// +build busybox

package handlers

import "net/http"

type MountFUSEFS struct {
}

func (p *MountFUSEFS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("fusefs not supported on busybox"))
}

type ListFUSEFS struct {
}

func (p *ListFUSEFS) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("fusefs not supported on busybox"))
}
