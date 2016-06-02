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
	c.JSON(w, 200, map[string]string{"derp": "sherp"})
	return nil
}
