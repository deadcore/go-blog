package post

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (m *postController) Delete(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")

	err := m.postDao.Delete(id); if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(200)
}