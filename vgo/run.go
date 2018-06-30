package vgo

import (
	"net/http"
	"fmt"
	"log"
	"vgo/router"
)

func Run(a string){
	fmt.Printf("%s is running",a)
	log.Fatal(http.ListenAndServe(a,router.Router))
}