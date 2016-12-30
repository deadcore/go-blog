package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PingController struct {}

func (m *PingController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}