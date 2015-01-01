package routes

import "github.com/tedsuo/rata"

const (
	Ping            = "PING"
	Env             = "ENV"
	Cwd             = "CWD"
	InstanceIndex   = "INDEX"
	StartedAt       = "STARTED_AT"
	ListExperiments = "LIST_EXPERIMENTS"
	Experiments     = "EXPERIMENTS"
	Hello           = "HELLO"
	Exit            = "EXIT"
	MakeTmpFile     = "MAKE_TMP_FILE"
	DeleteTmpFile   = "DELETE_TMP_FILE"
	Stick           = "STICK"
	Unstick         = "UNSTICK"
)

var Routes = rata.Routes{
	{Path: "/", Method: "GET", Name: Hello},
	{Path: "/ping", Method: "GET", Name: Ping},
	{Path: "/env", Method: "GET", Name: Env},
	{Path: "/cwd", Method: "GET", Name: Cwd},
	{Path: "/started-at", Method: "GET", Name: StartedAt},
	{Path: "/index", Method: "GET", Name: InstanceIndex},
	{Path: "/file/:filename", Method: "POST", Name: MakeTmpFile},
	{Path: "/file/:filename", Method: "DELETE", Name: DeleteTmpFile},
	{Path: "/exit/:code", Method: "POST", Name: Exit},
	{Path: "/experiments", Method: "GET", Name: ListExperiments},
	{Path: "/experiments/:experiment", Method: "GET", Name: Experiments},
	{Path: "/stick", Method: "GET", Name: Stick},
	{Path: "/unstick", Method: "GET", Name: Unstick},
}
