package handlers

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Curl struct {
}

func (p *Curl) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")
	if url == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing url"))
		return
	}

	resp, err := http.Get(url)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("faild to get URL: %s", err.Error())))
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("faild to read body: %s", err.Error())))
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}
