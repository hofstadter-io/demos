package main

import (
	"fmt"
	"os"

	"{{ .Datamodel.Module }}/pkg"
)

func main() {
	// temp
	if err := pkg.RunServer(); err != nil {
		fmt.Println(err)
		os.Exit(1)
  }

  if err := pkg.RootCmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}
