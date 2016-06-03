package web

import (
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/eknkc/amber"
	"github.com/julienschmidt/httprouter"
	"github.com/nii236/appreciator/controllers"
)

// Run starts the web server
func Run() {
	c := &controllers.HomeController{Render: amber.New()}

	r := httprouter.New()
	r.ServeFiles("/public/*filepath", http.Dir("public/"))
	r.Handler("GET", "/", c.Action(c.Index))
	r.Handler("POST", "/", c.Action(c.Create))

	n := negroni.Classic()
	n.UseHandler(r)

	err := http.ListenAndServe(":3001", n)
	if err != nil {
		log.Fatalln("Error!", err)
	}
}
