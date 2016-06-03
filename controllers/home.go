package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

// HomeController is the top level controller
type HomeController struct {
	BaseController
	*render.Render
}

// Index contains the index action for the home controller
func (c *HomeController) Index(w http.ResponseWriter, r *http.Request) error {
	c.HTML(w, http.StatusOK, "example", "world")
	return nil
}
