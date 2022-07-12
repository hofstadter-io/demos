type {{ camelT .name }} struct {
	// gorm.Model fields
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`

	// defined fields
	{{ range .Fields -}}
	{{ camelT .name }} {{ .type }}
	{{ end }}
	// relations
	{{ range .Reln -}}
	{{ if eq .type "OwnedBy" }}{{ camelT .name }}ID uint{{end}}
	{{ camelT .name }} {{ .goType }}
	{{ end }}
}
