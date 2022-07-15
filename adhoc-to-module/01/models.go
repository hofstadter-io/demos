package {{ .Datamodel.Name }}

{{ range .Datamodel.Models }}
type {{ .Name }} struct {
	// defined fields
	{{ range .Fields -}}
	{{ .Name }} {{ .Type }}
	{{ end }}

	// relations
	{{ range .Reln -}}
	{{ if eq .Type "OwnedBy" "HasOne" -}}
	{{ .Name }} {{ .Name }}
	{{ end -}}
	{{ if eq .Type "HasMany" "ManyToMany" -}}
	{{ .Name }} []{{ .Name }}
	{{ end -}}
	{{ end }}
}
{{ end }}
