{{ $ModelName := camelT .name }}
// {{ $ModelName }} handlers
func HandleCreate{{ $ModelName }}(c echo.Context) error {
	in := new({{ $ModelName }})
	if err := c.Bind(in); err != nil {
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
	if err := c.Bind(in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// {{ $ModelName }} ID from path `{{ kebab .name }}/:id`
	id := c.Param("id")
	// call library
	err := Update{{ $ModelName }}(id, in)
	if err != nil {
		return err
	}
	data := map[string]any{ "updated": id }
	return c.JSON(http.StatusOK, data)
}

func HandleDelete{{ $ModelName }}(c echo.Context) error {
	// {{ $ModelName }} ID from path `{{ kebab .name }}/:id`
	id := c.Param("id")
	// call library
	err := Delete{{ $ModelName }}(id)
	if err != nil {
		return err
	}
	// tell client
	data := map[string]any{ "deleted": id }
	return c.JSON(http.StatusOK, data)
}

