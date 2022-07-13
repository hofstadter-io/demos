package pkg

import (
	"net/http"
	"time"
	
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

{{/* it really does seem like turtles all the way down here */}}

{{ template "type/struct.go" . }}

{{ template "type/lib.go" . }}

{{ template "type/handler.go" . }}

{{ template "type/client.go" . }}

{{ template "type/command.go" . }}
