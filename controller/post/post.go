package post

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/deadcore/go-blog/model"
	"encoding/json"
	"bytes"
)

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
	buffer.WriteString("/post/")
	buffer.WriteString(post.Id)

	w.Header().Set("Location", buffer.String())
	w.WriteHeader(201)

}

