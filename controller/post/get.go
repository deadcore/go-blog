package post

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

func (m *postController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var id = p.ByName("id")

	var post, err = m.postDao.Get(id)

	if err != nil {
		http.Error(w, "", 404)
		return
	}

	json.NewEncoder(w).Encode(post)
}