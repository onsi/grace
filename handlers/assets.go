package handlers

import "text/template"

const stylesheet = `
`

var styledTemplate = template.Must(template.New("experiment").Parse(`
<html>
<head>
<style>
body {
    font-family: "helveticaneue-light";
    font-size: 16px;
    color: #333;
    margin:10px;
}
dt {
  color:#777;
}
</style>
</head>
<body>
{{.Body}}
</body>
</html>
`))

type Body struct {
	Body string
}
