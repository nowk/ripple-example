package main

import (
	"log"
	"net/http"

	"gopkg.in/labstack/echo.v1"
	"gopkg.in/nowk/ripple.v1"
)

// Api is a controller for our Api
type Api struct {
	Index echo.HandlerFunc `ripple:"GET /items"`
	Show  echo.HandlerFunc `ripple:"GET /items/:id"`
}

// Path mounts the Api controller group on /api/v1
func (Api) Path() string {
	return "/api/v1"
}

// IndexFunc returns JSON on GET /api/v1/items
func (Api) IndexFunc(ctx *echo.Context) error {
	return ctx.JSON(200, []interface{}{
		map[string]interface{}{
			"foo": "bar",
		},
	})
}

// ShowFunc returns JSON on /api/v1/items/:id
func (Api) ShowFunc(ctx *echo.Context) error {
	return ctx.JSON(200, map[string]interface{}{
		"id": ctx.Param("id"),
	})
}

func main() {
	// new echo
	mmux := echo.New()

	// mount our Api controller on the new echo
	ripple.Group(&Api{}, mmux)

	// start the server
	err := http.ListenAndServe(":4000", mmux)
	if err != nil {
		log.Fatalf("Server could not be started on :4000")
	}
}
