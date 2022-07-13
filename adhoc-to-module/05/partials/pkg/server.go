func RunServer() {
	if err := runServer(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

func runServer() error {
	// Setup the DB
	if err := initDB(); err != nil {
		return err
  }

	// Create the server
	e := echo.New()
	setupRouter(e)

	// get and format port
	port := "4242"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	port = ":" + port
	fmt.Println("{{ .Datamodel.Name }} listening on" + port)

	// run until we find the bottom turtle
	return e.Start(port)
}

func setupRouter(e *echo.Echo) {
	// setup recovery middleware
	e.Use(middleware.Recover())

	// setup logging middleware
	e.Use(middleware.Logger())

	e.GET("/internal/alive", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// Setup api routes
	{{- range .Datamodel.Models -}}
	{{- $ModelName := camelT .Name }}
	e.POST("/{{ kebab .Name }}", handleCreate{{ $ModelName }})
	e.GET("/{{ kebab .Name }}", handleList{{ $ModelName }})
	e.GET("/{{ kebab .Name }}/:id", handleGet{{ $ModelName }})
	e.PUT("/{{ kebab .Name }}/:id", handleUpdate{{ $ModelName }})
	e.DELETE("/{{ kebab .Name }}/:id", handleDelete{{ $ModelName }})
	{{- end }}
}

