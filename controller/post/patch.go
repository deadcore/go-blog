package post

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
	"github.com/imdario/mergo"
)

func (m *postController) Patch(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id := params.ByName("id")

	var updates map[string]interface{}
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&updates); if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	sourcePost, err := m.postDao.Get(id);

	if err := mergo.MapWithOverwrite(&sourcePost, updates); err != nil {
		panic(err)
	}

	m.postDao.Save(sourcePost)

	w.WriteHeader(200)
}