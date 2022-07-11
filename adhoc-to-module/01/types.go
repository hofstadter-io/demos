package demo

{{ range .Types }}
type {{ camelT .name }} struct {
	// defined fields
	{{ range .Fields -}}
	{{ camelT .name }} {{ .type }}
	{{ end }}

	// relations
	{{ range .Reln -}}
	{{ camelT .name }} {{ .goType }}
	{{ end }}
}
{{ end }}
