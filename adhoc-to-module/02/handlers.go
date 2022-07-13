package {{ .Datamodel.Name }}

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

// setup routes
func setupRoutes(e *echo.Echo) {
{{ range .Datamodel.Models }}
{{ $ModelName := camelT .Name }}
	e.POST("/{{ kebab .Name }}", handleCreate{{ $ModelName }})
	e.GET("/{{ kebab .Name }}", handleList{{ $ModelName }})
	e.GET("/{{ kebab .Name }}/:id", handleGet{{ $ModelName }})
	e.PUT("/{{ kebab .Name }}/:id", handleUpdate{{ $ModelName }})
	e.DELETE("/{{ kebab .Name }}/:id", handleDelete{{ $ModelName }})
{{ end }}
}

{{ range .Datamodel.Models }}
{{ $ModelName := camelT .Name }}
// {{ $ModelName }} handlers
func handleCreate{{ $ModelName }}(c echo.Context) error {
	in := new({{ .Name }})
	if err = c.Bind(in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ret, err := create{{ $ModelName }}(in)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}

func handleList{{ $ModelName }}(c echo.Context) error {
	ret, err := list{{ $ModelName }}()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}

func handleGet{{ $ModelName }}(c echo.Context) error {
	// {{ $ModelName }} ID from path `{{ kebab .Name }}/:id`
	id := c.Param("id")
	// call library
	ret, err := get{{ $ModelName }}ByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}

func handleUpdate{{ $ModelName }}(c echo.Context) error {
	in := new({{ $ModelName }})
	if err = c.Bind(in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// {{ $ModelName }} ID from path `{{ kebab .Name }}/:id`
	id := c.Param("id")
	// call library
	ret, err := update{{ $ModelName }}(id, in)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}

func handleDelete{{ $ModelName }}(c echo.Context) error {
	// {{ $ModelName }} ID from path `{{ kebab .Name }}/:id`
	id := c.Param("id")
	// call library
	ret, err := delete{{ $ModelName }}(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}
{{ end }}
