package {{ .Name }}

import (
	"gorm.io/gorm"
)

{{ range .Types }}
type {{ camelT .name }} struct {
	// ORM fields
	gorm.Model

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
