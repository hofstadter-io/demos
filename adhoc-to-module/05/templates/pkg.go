package pkg

import (
	"fmt"
	"net/http"
	"os"
	
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

{{ template "pkg/db.go" . }}

{{ template "pkg/server.go" . }}

{{ template "pkg/cli.go" . }}

