package actions

import (
    "net/http"
    
	"github.com/gobuffalo/buffalo"
)

// ServerpluginIndex default implementation.
func ServerpluginIndex(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("serverplugin/index.html"))
}

