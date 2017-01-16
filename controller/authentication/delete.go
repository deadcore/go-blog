package authentication

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (m *authenticationController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	token := params.ByName("token")

	if err := m.sessionService.Delete(token); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}