package main

import (
	"errors"
	"net/http"

	"path/filepath"

	"github.com/angelofallars/htmx-go"

	gowebly "github.com/gowebly/helpers"

	"github.com/gin-gonic/gin"
)

// indexViewHandler handles a view for the index page.
func indexViewHandler(c *gin.Context) {

	// Define paths to the user templates.
	indexPage := filepath.Join("templates", "pages", "index.html")

	// Parse user templates, using gowebly helper, or return error.
	tmpl, err := gowebly.ParseTemplates(indexPage)
	if err != nil {
		// If not, return HTTP 400 error.
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	// Execute (render) all templates or return error.
	if err := tmpl.Execute(c.Writer, nil); err != nil {
		// If not, return HTTP 500 error.
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

}

// showContentAPIHandler handles an API endpoint to show content.
func showContentAPIHandler(c *gin.Context) {
	// Check, if the current request has a 'HX-Request' header.
	// For more information, see https://htmx.org/docs/#request-headers
	if !htmx.IsHTMX(c.Request) {
		// If not, return HTTP 400 error.
		c.AbortWithError(http.StatusBadRequest, errors.New("non-htmx request"))
		return
	}

	// Write HTML content.
	c.Writer.Write([]byte("<p>🎉 Yes, <strong>htmx</strong> is ready to use! (<code>GET /api/hello-world</code>)</p>"))

	// Send htmx response.
	htmx.NewResponse().Write(c.Writer)
}
