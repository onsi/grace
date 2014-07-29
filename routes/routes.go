package routes

import "github.com/tedsuo/rata"

const (
	Ping          = "PING"
	Env           = "ENV"
	InstanceIndex = "INDEX"
)

var Routes = rata.Routes{
	{Path: "/ping", Method: "GET", Name: Ping},
	{Path: "/env", Method: "GET", Name: Env},
	{Path: "/index", Method: "GET", Name: InstanceIndex},
}
