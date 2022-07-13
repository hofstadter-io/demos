package {{ .Datamodel.Name }}

{{ range .Datamodel.Models }}
type {{ .Name }} struct {
	// defined fields
	{{ range .Fields -}}
	{{ .Name }} {{ .Type }}
	{{ end }}

	// relations
	{{ range .Reln -}}
	{{ .Name }} {{ .GoType }}
	{{ end }}
}
{{ end }}
