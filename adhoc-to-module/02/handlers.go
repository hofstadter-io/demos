package {{ camel .Datamodel.Name }}

import (
	"net/http"
	
	"github.com/labstack/echo/v4"
)

// routes
func setRoutes(e *echo.Echo) {
{{ range .Datamodel.Models }}
{{ $ModelName := camelT .name }}
	e.POST("/{{ kebab .name }}", HandleCreate{{ $ModelName }})
	e.GET("/{{ kebab .name }}", HandleList{{ $ModelName }})
	e.GET("/{{ kebab .name }}/:id", HandleGet{{ $ModelName }})
	e.PUT("/{{ kebab .name }}/:id", HandleUpdate{{ $ModelName }})
	e.DELETE("/{{ kebab .name }}/:id", HandleDelete{{ $ModelName }})
{{ end }}
}

{{ range .Datamodel.Models }}
{{ $ModelName := camelT .name }}
// {{ $ModelName }} handlers
func HandleCreate{{ $ModelName }}(c echo.Context) error {
	in := new({{ $ModelName }})
	if err = c.Bind(in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ret, err := Create{{ $ModelName }}(in)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}

func HandleList{{ $ModelName }}(c echo.Context) error {
	ret, err := List{{ $ModelName }}()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}

func HandleGet{{ $ModelName }}(c echo.Context) error {
	// {{ $ModelName }} ID from path `{{ kebab .name }}/:id`
	id := c.Param("id")
	// call library
	ret, err := Get{{ $ModelName }}ByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}

func HandleUpdate{{ $ModelName }}(c echo.Context) error {
	in := new({{ $ModelName }})
	if err = c.Bind(in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// {{ $ModelName }} ID from path `{{ kebab .name }}/:id`
	id := c.Param("id")
	// call library
	ret, err := Update{{ $ModelName }}(id, in)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}

func HandleDelete{{ $ModelName }}(c echo.Context) error {
	// {{ $ModelName }} ID from path `{{ kebab .name }}/:id`
	id := c.Param("id")
	// call library
	ret, err := Delete{{ $ModelName }}(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}
{{ end }}
