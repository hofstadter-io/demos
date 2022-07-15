type {{ .Name }} struct {
	// gorm fields
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`

	// defined fields
	{{ range .Fields -}}
	{{ .Name }} {{ .Type }} `{{ template "model/gorm.go" . }}`
	{{ end }}
	// relations
	{{ range .Reln -}}
	{{ if eq .Type "OwnedBy" }}{{ .Name }}ID uint{{end}}
	{{ if eq .Type "OwnedBy" "HasOne" -}}
	{{ .Name }} {{ .Name }}
	{{ end -}}
	{{ if eq .Type "HasMany" "ManyToMany" -}}
	{{ .Name }} []{{ .Name }}
	{{ end -}}
	{{ end }}
}
