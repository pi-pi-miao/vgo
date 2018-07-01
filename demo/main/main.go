package main

import (
	_ "github.com/PyreneGitHub/vgo/router"
	"github.com/PyreneGitHub/vgo/vgo"
)

func main() {
	ConfigAll()
	vgo.Run(":8080")
}
