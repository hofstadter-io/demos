package {{ .Datamodel.Name }}

{{ range .Datamodel.Models }}
type {{ .Name }} struct {
	// defined fields
	{{ range .Fields -}}
	{{ .Name }} {{ .Type }}
	{{ end }}

	// relations
	{{ range .Reln -}}
	{{ if eq .Type "OwnedBy" }}{{ .Name }}ID uint{{end}}
	{{ .Name }} {{ .GoType }}
	{{ end }}
}
{{ end }}
