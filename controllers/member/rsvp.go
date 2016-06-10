package member

import (
	"net/http"

	"github.com/eknkc/amber"
	"github.com/julienschmidt/httprouter"
	"github.com/nii236/appreciator/controllers"
)

// RSVPController is the top level controller
type RSVPController struct {
	controllers.BaseController
	Render *amber.Compiler
}

// Index contains the index action for the RSVP controller
func (c *RSVPController) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	err := c.Render.ParseFile("./views/member/rsvp/login.amber")
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

// RSVP contains a form which guests can POST to confirm their RSVP
func (c *RSVPController) RSVP(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	err := c.Render.ParseFile("./views/member/rsvp/rsvp.amber")
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

// Complete contains the action after a successful RSVP for the RSVP controller
func (c *RSVPController) Complete(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	err := c.Render.ParseFile("./views/member/rsvp/complete.amber")
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

// Login contains the Login action for the RSVP controller
func (c *RSVPController) Login(w http.ResponseWriter, r *http.Request, params httprouter.Params) error {
	r.ParseForm()
	http.Redirect(w, r, "/rsvp/do", http.StatusFound)
	return nil
}

func GetGuest() string {
	return ""
}
