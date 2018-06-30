package main

import(
	"vgo/vgo"
	_ "vgo/router"
)

func main(){
	ConfigAll()
	vgo.Run(":8080")
}