package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// BaseController as a base controller for the web app
type BaseController struct{}

// Action helps with error handling in a controller
type Action func(w http.ResponseWriter, r *http.Request, params httprouter.Params) error

// Action is a base action that returns a Handler
func (c *BaseController) Action(a Action) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if err := a(w, r, params); err != nil {
			http.Error(w, err.Error(), 500)
		}
	}
}
