package {{ .Datamodel.Name }}

import (
	"gorm.io/gorm"
	"gorm.io/gorm/sqlite"
)

var db gorm.DB

func InitDB() (err error) {
	// Setup the database connection
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema for each Model
	{{ range .Datamodel.Models -}}
	db.AutoMigrate(&{{ .Name }}{})
	{{ end }}

	return nil
}
