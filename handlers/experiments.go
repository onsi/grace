package handlers

import (
	"html/template"
	"net/http"

	"github.com/onsi/grace/handlers/experiments"
)

type ListExperiments struct {
}

func (p *ListExperiments) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	experimentTemplate.Execute(w, experiments.Experiments)
}

var experimentTemplate = template.Must(template.New("experiment").Parse(`
<html>
<head>
</head>
<body>
<table>
{{range $experimentPath, $experiment := .}}
<tr>
  <td><a href="/experiments/{{$experimentPath}}">{{$experiment.Name}}</td>
  <td>{{$experiment.Description}}</td>
</tr>
{{end}}
</table>
</body>
</html>
`))

type Experiments struct {
}

func (p *Experiments) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	requestedExperiment := r.URL.Query().Get(":experiment")
	experiment, ok := experiments.Experiments[requestedExperiment]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	experiment.Handler(w, r)
}
