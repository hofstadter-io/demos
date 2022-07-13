var db *gorm.DB

func initDB() (err error) {
	// Create db connection
	db, err = gorm.Open(sqlite.Open("{{ .Name }}.db"), &gorm.Config{})
  if err != nil {
		return fmt.Errorf("failed to connect database:\n%s", err)
  }

  // Migrate the schema
	{{ range .Datamodel.Models -}}
  err = db.AutoMigrate(&{{ .Name }}{})
  if err != nil {
		return fmt.Errorf("failed to migrate database:\n%s", err)
  }
  {{ end }}

	return nil
}
