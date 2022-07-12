package pkg

import (
	"net/http"
	"time"
	
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

{{/* it really does seem like turtles all the way down here */}}

{{ template "type.go" . }}

{{ template "lib.go" . }}

{{ template "handler.go" . }}
