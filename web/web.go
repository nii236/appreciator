package web

import (
	"log"
	"net/http"

	"github.com/unrolled/render"

	"github.com/gorilla/mux"
	"github.com/nii236/appreciator/controllers"
)

// Run starts the web server
func Run() {
	c := &controllers.HomeController{Render: render.New(render.Options{})}
	r := mux.NewRouter()
	r.Handle("/", c.Action(c.Index)).Methods("GET")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("Error!", err)
	}
}
