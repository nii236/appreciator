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
	r.GET("/", c.Action(c.Index))
	r.POST("/", c.Action(c.Create))

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	n.UseHandler(r)

	err := http.ListenAndServe(":3001", n)
	if err != nil {
		log.Fatalln("Error!", err)
	}
}
