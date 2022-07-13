package pkg

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() (err error) {
	// Create db connection
	db, err = gorm.Open(sqlite.Open("{{ .Datamodel.Name }}.db"), &gorm.Config{})
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

func SetupRouter(e *echo.Echo) {
	// setup recovery middleware
	e.Use(middleware.Recover())

	// setup logging middleware
	e.Use(middleware.Logger())

	e.GET("/internal/alive", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// Setup api routes
	{{- range .Models -}}
	{{- $ModelName := camelT .Name }}
	e.POST("/{{ kebab .Name }}", handleCreate{{ $ModelName }})
	e.GET("/{{ kebab .Name }}", handleList{{ $ModelName }})
	e.GET("/{{ kebab .Name }}/:id", handleGet{{ $ModelName }})
	e.PUT("/{{ kebab .Name }}/:id", HandleUpdate{{ $ModelName }})
	e.DELETE("/{{ kebab .Name }}/:id", handleDelete{{ $ModelName }})
	{{- end }}
}
