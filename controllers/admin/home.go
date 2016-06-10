package admin

import (
	"net/http"

	"github.com/eknkc/amber"
	"github.com/julienschmidt/httprouter"
	"github.com/nii236/appreciator/controllers"
)

// HomeController is the top level controller
type HomeController struct {
	controllers.BaseController
	Render *amber.Compiler
}

// Index contains the index action for the home controller
func (c *HomeController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	return nil
}
