package controller

import (
	"Go-Chi/Model"
	"Go-Chi/driver"
	"Go-Chi/services"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type Post struct {
	repo services.PostService
}

func(service *Post) AllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts,_ = service.repo.AllPosts()
	json.NewEncoder(w).Encode(posts)

	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
}

//func AllPosts(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var posts,_ = services.AllPosts()
//	json.NewEncoder(w).Encode(posts)
//
//	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
//}

//func DetailPost(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	id := chi.URLParam(r, "id")
//	var post = services.DetailPost(id)
//	json.NewEncoder(w).Encode(post)
//	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
//}
//
//func CreatePost(w http.ResponseWriter, r *http.Request) {
//	var post Model.Post
//	json.NewDecoder(r.Body).Decode(&post)
//	services.CreatePost(post)
//
//	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
//}

// UpdatePost update a  spesific post
func UpdatePost(w http.ResponseWriter, r *http.Request) {
	var post Model.Post
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&post)

	query, err := driver.DBConn().Prepare("Update posts set title=?, content=? where id=?")
	catch(err)
	_, er := query.Exec(post.Title, post.Content, id)
	catch(er)

	defer query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})

}

// DeletePost remove a spesific post
func DeletePost(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query, err := driver.DBConn().Prepare("delete from posts where id=?")
	catch(err)
	_, er := query.Exec(id)
	catch(er)
	query.Close()

	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}