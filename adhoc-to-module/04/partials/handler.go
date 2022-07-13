{{ $ModelName := camelT .Name }}
// {{ $ModelName }} handlers
func handleCreate{{ $ModelName }}(c echo.Context) error {
	in := new({{ $ModelName }})
	if err := c.Bind(in); err != nil {
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
	if err := c.Bind(in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// {{ $ModelName }} ID from path `{{ kebab .Name }}/:id`
	id := c.Param("id")
	// call library
	err := update{{ $ModelName }}(id, in)
	if err != nil {
		return err
	}
	data := map[string]any{ "updated": id }
	return c.JSON(http.StatusOK, data)
}

func handleDelete{{ $ModelName }}(c echo.Context) error {
	// {{ $ModelName }} ID from path `{{ kebab .Name }}/:id`
	id := c.Param("id")
	// call library
	err := delete{{ $ModelName }}(id)
	if err != nil {
		return err
	}
	// tell client
	data := map[string]any{ "deleted": id }
	return c.JSON(http.StatusOK, data)
}

