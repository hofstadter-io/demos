
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

func RunServer() error {
	// Setup the DB
	if err := InitDB(); err != nil {
		fmt.Println(err)
		os.Exit(1)
  }

	// Create the server
	e := echo.New()
	SetupRouter(e)

	// get and format port
	port := "4242"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	port = ":" + port
	fmt.Println("{{ .Datamodel.Name }} listening on" + port)

	// run until we find the bottom turtle
	e.Logger.Fatal(e.Start(port))

	return nil
}
