package experiments

import "net/http"

type Experiment struct {
	Name        string
	Description string
	Handler     http.HandlerFunc
}

var Experiments = map[string]Experiment{}
