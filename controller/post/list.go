package post

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

func (m *postController) List(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	results, err := m.postDao.FindAll()
	if err != nil {
		panic(err)
	}
	json.NewEncoder(w).Encode(results)
}