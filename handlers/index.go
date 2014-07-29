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
		w.Write([]byte("Couldn't find VCAP_APPLICATION"))
		w.WriteHeader(http.StatusInternalServerError)
	}

	var decodedVcap struct {
		InstanceIndex int `json:"instance_index"`
	}

	err := json.Unmarshal([]byte(vcap), &decodedVcap)
	if err != nil {
		w.Write([]byte("Couldn't parse VCAP_APPLICATION"))
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Write([]byte(fmt.Sprintf("%d", decodedVcap.InstanceIndex)))
	w.WriteHeader(http.StatusOK)
}
