package router

import (
	"github.com/PyreneGitHub/vgo/demo/controller"
	"github.com/PyreneGitHub/vgo/router"
)

func init() {
	router.Router.GET("/", controller.T)
}
