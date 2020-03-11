package repo_iplm

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"fmt"

	"Go-Chi/Model"
	//"encoding/json"
	//"github.com/go-chi/chi"
	//"net/http"
)

type PostRepoIplm struct {
	Db *sql.DB
}

func (u *PostRepoIplm) AllPosts() ([]Model.Post, error) {
	posts := make([]Model.Post, 0)

	rows, err := u.Db.Query("SELECT * FROM posts")
	if err != nil {
		return posts, err
	}
	defer rows.Close()
	for rows.Next() {
		post := Model.Post{}
		// SELECT id, name, gender, email FROM public.users;
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			break
		}

		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func (*PostRepoIplm) DetailPost(id int) (Model.Post, error) {
	panic("implement me")
}

func (*PostRepoIplm) CreatePost(u Model.Post) (error) {
	panic("implement me")
}

func (*PostRepoIplm) UpdatePost(id int) (error) {
	panic("implement me")
}

func (*PostRepoIplm) DeletePost(id int) (error) {
	panic("implement me")
}

//func NewPostRepo(db *sql.DB) repository.PostRepo {
//	return &PostRepoIplm {
//		Db : db,
//	}
//}

//func AllPosts(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var post Model.Post
//	json.NewDecoder(r.Body).Decode(&post)
//
//	query, err := driver.DBConn().Query("SELECT * FROM posts")
//	catch(err)
//
//	var posts []Model.Post
//	for query.Next(){
//		var id int
//		var title, content string
//
//		err = query.Scan(&id, &title, &content)
//		if err != nil {
//			panic(err.Error())
//		}
//
//		post.ID = id
//		post.Title = title
//		post.Content = content
//		posts = append(posts, post)
//	}
//	defer query.Close()
//	json.NewEncoder(w).Encode(posts)
//
//	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
//}
//
//func DetailPost(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	var post Model.Post
//	json.NewDecoder(r.Body).Decode(&post)
//
//	query, err := driver.DBConn().Query("SELECT * FROM posts")
//	catch(err)
//
//	for query.Next(){
//		var id int
//		var title, content string
//
//		err = query.Scan(&id, &title, &content)
//		if err != nil {
//			panic(err.Error())
//		}
//
//		post.ID = id
//		post.Title = title
//		post.Content = content
//	}
//	defer query.Close()
//	json.NewEncoder(w).Encode(post)
//
//	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
//}
//
//func CreatePost(w http.ResponseWriter, r *http.Request) {
//	var post Model.Post
//	json.NewDecoder(r.Body).Decode(&post)
//
//	query, err := driver.DBConn().Prepare("Insert posts SET title=?, content=?")
//	catch(err)
//
//	_, er := query.Exec(post.Title, post.Content)
//	catch(er)
//	defer query.Close()
//
//	respondwithJSON(w, http.StatusCreated, map[string]string{"message": "successfully created"})
//}
//
//// UpdatePost update a  spesific post
//func UpdatePost(w http.ResponseWriter, r *http.Request) {
//	var post Model.Post
//	id := chi.URLParam(r, "id")
//	json.NewDecoder(r.Body).Decode(&post)
//
//	query, err := driver.DBConn().Prepare("Update posts set title=?, content=? where id=?")
//	catch(err)
//	_, er := query.Exec(post.Title, post.Content, id)
//	catch(er)
//
//	defer query.Close()
//
//	respondwithJSON(w, http.StatusOK, map[string]string{"message": "update successfully"})
//
//}
//
//// DeletePost remove a spesific post
//func DeletePost(w http.ResponseWriter, r *http.Request) {
//	id := chi.URLParam(r, "id")
//
//	query, err := driver.DBConn().Prepare("delete from posts where id=?")
//	catch(err)
//	_, er := query.Exec(id)
//	catch(er)
//	query.Close()
//
//	respondwithJSON(w, http.StatusOK, map[string]string{"message": "successfully deleted"})
//}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}

//// respondwithError return error message
//func respondWithError(w http.ResponseWriter, code int, msg string) {
//	respondwithJSON(w, code, map[string]string{"message": msg})
//}
//
// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	fmt.Println(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}