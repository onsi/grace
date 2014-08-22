package vcap_application_parser

import (
	"encoding/json"
	"errors"
	"os"
)

func GetIndex() (int, error) {
	vcap := os.Getenv("VCAP_APPLICATION")
	if vcap == "" {
		return 0, errors.New("couldn't find VCAP_APPLICATION")
	}

	var decodedVcap struct {
		InstanceIndex int `json:"instance_index"`
	}

	err := json.Unmarshal([]byte(vcap), &decodedVcap)
	if err != nil {
		return 0, errors.New("couldn't parse VCAP_APPLICATION")
	}

	return decodedVcap.InstanceIndex, nil
}
