package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"

	"{{ .Module }}/pkg"
)

func main() {
	// Setup the DB
	if err := pkg.InitDB(); err != nil {
		fmt.Println(err)
		os.Exit(1)
  }

	// Create the server
	e := echo.New()
	pkg.SetupRouter(e)

	// get and format port
	port := "4242"
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	port = ":" + port
	fmt.Println("listening on" + port)

	// run until we find the bottom turtle
	e.Logger.Fatal(e.Start(port))
}
