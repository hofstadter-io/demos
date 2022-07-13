package pkg

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

{{/* it really does seem like turtles all the way down here */}}

{{ template "model/struct.go" .Model }}

{{ template "model/lib.go" .Model }}

{{ template "model/handler.go" .Model }}

{{ if .Config.Client.go.enabled }}
{{ template "model/client.go" .Model }}

{{ if .Config.Cli.enabled }}
{{ template "model/command.go" .Model }}
{{ end }}

{{ end }}

