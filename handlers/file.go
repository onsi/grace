package handlers

import (
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

type MakeTmpFile struct {
}

var fileRe = regexp.MustCompile(`[A-Za-z0-9_- ]+`)

func (p *MakeTmpFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	fileName := r.URL.Query().Get(":filename")
	if !fileRe.MatchString(fileName) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid file"))
		return
	}
	f, err := os.Create("/tmp/" + fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	_, err = f.Write(content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}

type DeleteTmpFile struct {
}

func (p *DeleteTmpFile) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get(":filename")
	if !fileRe.MatchString(fileName) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid file"))
		return
	}
	err := os.RemoveAll("/tmp/" + fileName)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
}
