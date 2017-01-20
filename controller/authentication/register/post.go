package register

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"bytes"
)

type credentials struct {
	Email    string
	Password string
}

func (m *registrationController) Post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
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

	user := m.userService.Register(creds.Email, creds.Password)

	var buffer bytes.Buffer
	buffer.WriteString("/users/")
	buffer.WriteString(user.Id)

	w.Header().Set("Location", buffer.String())

	json.NewEncoder(w).Encode(user)
	w.WriteHeader(201)
}