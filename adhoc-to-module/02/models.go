package {{ .Datamodel.Name }}

import (
	"gorm.io/gorm"
)

{{ range .Datamodel.Models }}
type {{ .Name }} struct {
	// ORM fields
	gorm.Model

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
