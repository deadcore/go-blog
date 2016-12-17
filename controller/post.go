package controller

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"api.khazix.co.uk/dao"
	"encoding/json"
	"api.khazix.co.uk/model"
	"bytes"
	"strconv"
)

type PostController struct {
	postDao dao.PostDao
}

func (m *PostController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var id, _ = strconv.ParseUint(p.ByName("id"), 10, 64)
	var post, err = m.postDao.Get(id)

	if err != nil {
		http.Error(w, "", 404)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(post)
}

func (m *PostController) Post(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var post model.Post
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	post = m.postDao.Save(post)

	var buffer bytes.Buffer
	var stringId = strconv.FormatUint(post.Id, 10)
	buffer.WriteString("/post/")
	buffer.WriteString(stringId)

	w.Header().Set("Location", buffer.String())
	w.WriteHeader(201)

}