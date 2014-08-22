package handlers

import "text/template"

var styledTemplate = template.Must(template.New("experiment").Parse(`
<html>
<head>
<style>
body {
    font-family: "helveticaneue-light";
    font-size: 16px;
    color: #333;
    position: absolute;
    margin:0;
    width:100%;
    height:100%;
}

dt {
  color:#777;
}

.envs {
  margin:10px
}

.hello {
  position:absolute;
  top:0;
  height:120px;
  left:0;
  right:0;
  text-align:center;
  font-size:80px;
  font-weight:bold;
  line-height:120px;
  color: rgb(0,139,185);
}

.my-index {
  position:absolute;
  top:120px;
  height:30px;
  color: #333;
  font-size:30px;
  line-height:30px;
  left:0;
  right:0;
  text-align:center;
  color: rgb(0,151,198)
}

.index {
  position:absolute;
  top:176px;
  height:120px;
  left:0;
  right:0;
  color: #fff;
  font-size: 80px;
  line-height: 120px;
  background-color:rgb(36,184,235);
  text-align:center;
}

.mid-color {
  position:absolute;
  top:296px;
  height:120px;
  left:0;
  right:0;
  background-color: rgb(0,151,198)
}
.bottom-color {
  position:absolute;
  top:416px;
  bottom:0;
  left:0;
  right:0;
  background-color: rgb(0,139,185);
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
