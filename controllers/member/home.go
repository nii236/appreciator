package member

import (
	"net/http"

	log "github.com/Sirupsen/logrus"
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
	err := c.Render.ParseFile("./views/member/form.amber")
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
	log.Debugln("to1:", r.FormValue("to1"))
	log.Debugln("to2:", r.FormValue("to2"))
	log.Debugln("to3:", r.FormValue("to3"))
	log.Debugln("attended:", r.FormValue("attended"))
	log.Debugln("gift:", r.FormValue("gift"))
	log.Debugln("additional:", r.FormValue("additional"))
	log.Debugln("from1:", r.FormValue("from1"))
	log.Debugln("from2:", r.FormValue("from2"))
	http.Redirect(w, r, "/member", http.StatusFound)
	return nil
}
