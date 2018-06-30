package controller

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
)


func T(w http.ResponseWriter, r *http.Request,c httprouter.Params) {
	_,cancel:=context.WithCancel(context.Background())
	defer cancel()
	fmt.Fprint(w, "Welcome!\n")
}