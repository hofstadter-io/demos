{{ $TypeName := camelT .name }}
// {{ $TypeName }} handlers
func HandleCreate{{ $TypeName }}(c echo.Context) error {
	in := new({{ $TypeName }})
	if err := c.Bind(in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	ret, err := Create{{ $TypeName }}(in)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}

func HandleList{{ $TypeName }}(c echo.Context) error {
	ret, err := List{{ $TypeName }}()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}

func HandleGet{{ $TypeName }}(c echo.Context) error {
	// {{ $TypeName }} ID from path `{{ kebab .name }}/:id`
	id := c.Param("id")
	// call library
	ret, err := Get{{ $TypeName }}ByID(id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, ret)
}

func HandleUpdate{{ $TypeName }}(c echo.Context) error {
	in := new({{ $TypeName }})
	if err := c.Bind(in); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	// {{ $TypeName }} ID from path `{{ kebab .name }}/:id`
	id := c.Param("id")
	// call library
	err := Update{{ $TypeName }}(id, in)
	if err != nil {
		return err
	}
	data := map[string]any{ "updated": id }
	return c.JSON(http.StatusOK, data)
}

func HandleDelete{{ $TypeName }}(c echo.Context) error {
	// {{ $TypeName }} ID from path `{{ kebab .name }}/:id`
	id := c.Param("id")
	// call library
	err := Delete{{ $TypeName }}(id)
	if err != nil {
		return err
	}
	// tell client
	data := map[string]any{ "deleted": id }
	return c.JSON(http.StatusOK, data)
}

