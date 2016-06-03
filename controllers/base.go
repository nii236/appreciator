package controllers

import "net/http"

// BaseController as a base controller for the web app
type BaseController struct{}

// Action helps with error handling in a controller
type Action func(w http.ResponseWriter, r *http.Request) error

// Action is a base action that returns a Handler
func (c *BaseController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := a(w, r); err != nil {
			http.Error(w, err.Error(), 500)
		}
	})
}
