package web

import (
	"net/http"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"

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
	memberRSVP := &member.RSVPController{Render: amber.New()}

	r := httprouter.New()
	r.ServeFiles("/public/*filepath", http.Dir("public/"))

	r.GET("/admin", adminHome.Action(adminHome.Index))
	r.GET("/member", memberHome.Action(memberHome.Index))
	r.POST("/member", memberHome.Action(memberHome.Create))

	r.GET("/rsvp", memberRSVP.Action(memberRSVP.Index))
	r.GET("/rsvp/do", memberRSVP.Action(memberRSVP.RSVP))
	r.POST("/rsvp/complete", memberRSVP.Action(memberRSVP.Complete))
	r.POST("/rsvp", memberRSVP.Action(memberRSVP.Login))

	n := negroni.New()
	n.Use(negroni.HandlerFunc(loggerMiddleware))
	n.Use(negroni.NewRecovery())
	n.UseHandler(r)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":3000"
	}
	log.Infoln("Starting server on", PORT)
	err := http.ListenAndServe(PORT, n)
	if err != nil {
		log.Fatalln("Error!", err)
	}
}

func loggerMiddleware(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	then := time.Now()
	log.Infoln("Started", r.Method, r.URL)
	next(rw, r)
	log.Infoln("Finished", r.Method, r.URL, "in", time.Since(then))
}
