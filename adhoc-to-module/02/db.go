package {{ camel .Datamodel.Name }}

import (
	"gorm.io/gorm"
	"gorm.io/gorm/sqlite"
)

var db gorm.DB

func InitDB() (err error) {
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
  if err != nil {
    panic("failed to connect database")
  }

  // Migrate the schema
	{{ range .Datamodel.Models -}}
  db.AutoMigrate(&{{ camelT .name }}{})
  {{ end }}

	return nil
}
