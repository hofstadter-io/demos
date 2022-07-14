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
// comment at the edge of code gen loop
// help prevent git like merge conflicts
// with custom code in diff3 mode
