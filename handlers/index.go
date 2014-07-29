package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type InstanceIndex struct {
}

func (p *InstanceIndex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vcap := os.Getenv("VCAP_APPLICATION")
	if vcap == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Couldn't find VCAP_APPLICATION"))
	}

	var decodedVcap struct {
		InstanceIndex int `json:"instance_index"`
	}

	err := json.Unmarshal([]byte(vcap), &decodedVcap)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Couldn't parse VCAP_APPLICATION"))
	}

	w.Write([]byte(fmt.Sprintf("%d", decodedVcap.InstanceIndex)))
}
