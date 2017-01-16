package authentication

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

type credentials struct {
	Email    string
	Password string
}

func (m *authenticationController) Post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var creds credentials
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	session, err := m.authenticationService.Authenticate(creds.Email, creds.Password)

	if err != nil {
		http.Error(w, "", 401)
		return
	}

	w.Header().Set("x-session-token", session.Token)
	json.NewEncoder(w).Encode(session)
}