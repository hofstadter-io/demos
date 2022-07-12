package pkg

import (
	"net/http"
	"time"
	
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

{{ template "partials/struct.go" . }}

{{ template "partials/lib.go" . }}

{{ template "partials/handler.go" . }}
