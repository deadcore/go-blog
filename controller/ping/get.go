package ping

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (m *PingController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}