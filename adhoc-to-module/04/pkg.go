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
	db, err = gorm.Open(sqlite.Open("{{ .Name }}.db"), &gorm.Config{})
  if err != nil {
		return fmt.Errorf("failed to connect database:\n%s", err)
  }

  // Migrate the schema
	{{ range .Datamodel.Models -}}
  err = db.AutoMigrate(&{{ camelT .name }}{})
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
	{{- range .Datamodel.Models -}}
	{{- $ModelName := camelT .name }}
	e.POST("/{{ kebab .name }}", HandleCreate{{ $ModelName }})
	e.GET("/{{ kebab .name }}", HandleList{{ $ModelName }})
	e.GET("/{{ kebab .name }}/:id", HandleGet{{ $ModelName }})
	e.PUT("/{{ kebab .name }}/:id", HandleUpdate{{ $ModelName }})
	e.DELETE("/{{ kebab .name }}/:id", HandleDelete{{ $ModelName }})
	{{- end }}
}
