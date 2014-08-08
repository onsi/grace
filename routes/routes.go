package routes

import "github.com/tedsuo/rata"

const (
	Ping            = "PING"
	Env             = "ENV"
	InstanceIndex   = "INDEX"
	ListExperiments = "LIST_EXPERIMENTS"
	Experiments     = "EXPERIMENTS"
)

var Routes = rata.Routes{
	{Path: "/ping", Method: "GET", Name: Ping},
	{Path: "/env", Method: "GET", Name: Env},
	{Path: "/index", Method: "GET", Name: InstanceIndex},
	{Path: "/experiments", Method: "GET", Name: ListExperiments},
	{Path: "/experiments/:experiment", Method: "GET", Name: Experiments},
}
