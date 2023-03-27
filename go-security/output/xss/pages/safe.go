package pages

var TemplateSafe = `<!DOCTYPE html>
<html>
	<head>
		<title>{{.Name}}'s Messages</title>
	</head>
	<body>
	<h1>{{.Name}}'s Messages</h1>
	<p>
		You have {{len .Messages}} messages.
	</p>
	{{range .Messages}}
	<p>[{{.Time.Format "2006-01-02T15:04"}} {{.From}}] {{.Content}}</p><hr />
	{{end}}
	</body>
</html>
`
