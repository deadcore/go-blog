package ping

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (m *pingController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
}