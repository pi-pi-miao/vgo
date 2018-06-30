package router

import(
	"vgo/router"
	"go_dev/day19/httprouter/controller"
)

func init(){
	router.Router.GET("/",controller.T)
}
