package router

import(
	"vgo/router"
	"vgo/demo/controller"
)

func init(){
	router.Router.GET("/",controller.T)
}
