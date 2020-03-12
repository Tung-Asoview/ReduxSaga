package repository

import (
	"Go-Chi/Model"
	"Go-Chi/services"
	"Go-Chi/driver"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

type Service struct {
	Db *sql.DB
}

func NewPostRepo(db *sql.DB) services.PostService {
	return &Service {
		Db: db,
	}
}

func (s *Service) AllPosts() ([]Model.Post, error){
	var post Model.Post

	query, err := driver.DBConn().Query("SELECT * FROM posts")
	catch(err)

	var posts []Model.Post
	for query.Next(){
		var id int
		var title, content string

		err = query.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}

		post.ID = id
		post.Title = title
		post.Content = content
		posts = append(posts, post)
	}

	return posts, err
}

//func (s *Service) DetailPost(id int) (*Model.Post, error){
//
//}
//
//func (s *Service) CreatePost(u *Model.Post) (error){
//
//}
//
//func (s *Service) UpdatePost(id int) (error){
//
//}
//
//func (s *Service) DeletePost(id int) (error){
//
//}


//func (u *PostRepo) AllPosts() ([]Model.Post, error) {
//	posts := make([]Model.Post, 0)
//
//	rows, err := u.Db.Query("SELECT * FROM posts")
//	if err != nil {
//		return posts, err
//	}
//	defer rows.Close()
//	for rows.Next() {
//		post := Model.Post{}
//		// SELECT id, name, gender, email FROM public.users;
//		err := rows.Scan(&post.ID, &post.Title, &post.Content)
//		if err != nil {
//			break
//		}
//
//		posts = append(posts, post)
//	}
//
//	err = rows.Err()
//	if err != nil {
//		return posts, err
//	}
//
//	return posts, nil
//}

func AllPosts() ([]Model.Post, error) {
	var post Model.Post

	query, err := driver.DBConn().Query("SELECT * FROM posts")
	catch(err)

	var posts []Model.Post
	for query.Next(){
		var id int
		var title, content string

		err = query.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}

		post.ID = id
		post.Title = title
		post.Content = content
		posts = append(posts, post)
	}

	return posts, err
}

func DetailPost(id string)(Model.Post) {
	var post Model.Post

	query, err := driver.DBConn().Query("SELECT * FROM posts WHERE id=?", id)
	catch(err)

	for query.Next(){
		var id int
		var title, content string

		err = query.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}

		post.ID = id
		post.Title = title
		post.Content = content
	}

	return post
}

func CreatePost(p Model.Post) {

	query, err := driver.DBConn().Prepare("Insert posts SET title=?, content=?")
	catch(err)

	_, er := query.Exec(p.Title, p.Content)
	catch(er)
	defer query.Close()
}

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

// respondwithError return error message
//func respondWithError(w http.ResponseWriter, code int, msg string) {
//	respondwithJSON(w, code, map[string]string{"message": msg})
//}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}