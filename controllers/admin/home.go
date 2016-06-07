package admin

import (
	"log"
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
	err := c.Render.ParseFile("./views/basic.amber")
	if err != nil {
		return err
	}
	tpl, err := c.Render.Compile()
	if err != nil {
		return err
	}
	if err := tpl.Execute(w, nil); err != nil {
		return err
	}
	return nil
}

// Create contains the create action for the home controller
func (c *HomeController) Create(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	r.ParseForm()
	log.Println(r.Form)
	http.Redirect(w, r, "/", http.StatusFound)
	return nil
}
