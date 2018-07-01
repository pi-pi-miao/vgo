package vgo

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PyreneGitHub/vgo/router"
)

func Run(a string) {
	fmt.Printf("%s is running", a)
	log.Fatal(http.ListenAndServe(a, router.Router))
}
