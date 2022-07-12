package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"

	"{{ .Datamodel.Module }}/pkg"
)

func main() {
  if err := pkg.RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
