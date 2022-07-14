var host string

func init() {
	host = "http://localhost:4242"
	// req.DevMode()
}

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

	// Start server
	go func() {
		fmt.Println("{{ .Datamodel.Name }} listening on" + port)
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			e.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds. 
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
	fmt.Println("{{ .Datamodel.Name }} exiting")

	return nil
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

