package web

import (
	"log"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/eknkc/amber"
	"github.com/julienschmidt/httprouter"
	"github.com/nii236/appreciator/controllers/admin"
	"github.com/nii236/appreciator/controllers/member"
)

// Run starts the web server
func Run() {
	adminHome := &admin.HomeController{Render: amber.New()}
	memberHome := &member.HomeController{Render: amber.New()}

	r := httprouter.New()
	r.ServeFiles("/public/*filepath", http.Dir("public/"))

	r.GET("/admin", adminHome.Action(adminHome.Index))
	r.POST("/admin", adminHome.Action(adminHome.Create))

	r.GET("/member", memberHome.Action(memberHome.Index))
	r.POST("/member", memberHome.Action(memberHome.Create))

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	n.UseHandler(r)
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3000"
	}
	log.Println("Starting server on", PORT)
	err := http.ListenAndServe(PORT, n)
	if err != nil {
		log.Fatalln("Error!", err)
	}
}
