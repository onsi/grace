package handlers

import (
	"fmt"
	"net/http"

	"github.com/onsi/grace/vcap_application_parser"
)

type InstanceIndex struct {
}

func (p *InstanceIndex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	index, err := vcap_application_parser.GetIndex()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write([]byte(fmt.Sprintf("%d", index)))
}
