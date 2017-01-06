package post

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

func (m *PostController) List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	json.NewEncoder(w).Encode(m.PostDao.FindAll())
}