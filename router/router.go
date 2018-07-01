package router

import (
	"github.com/julienschmidt/httprouter"
)

func RouterMethod() *httprouter.Router {
	routered := httprouter.New()
	Router = routered
	return Router
}
