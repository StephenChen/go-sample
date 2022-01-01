package router

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Router() {
	r := httprouter.New()
	r.PUT("/usr/installations/:installation_id/repositories/:reposit", Hello)
	r.GET("/user/status", Hello)
	r.GET("/user/search", Hello)
	r.GET("/user/support", Hello)
	r.GET("/user/super", Hello)
	r.GET("/user/susan", Hello)
	r.GET("/user/s/:id", Hello)

}

func Hello(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

}
