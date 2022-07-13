type {{ .Name }} struct {
	// gorm.Model fields
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`

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
