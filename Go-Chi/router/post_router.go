package router

import (
	"fmt"
	"Go-Chi/repository/repo_iplm"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi"
	"net/http"
	"time"
)

var router *chi.Mux
func Routers() *chi.Mux {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Get("/posts", repo_iplm.AllPosts)
	//router.Get("/posts/{id}", repository.DetailPost)
	//router.Post("/posts", repository.CreatePost)
	//router.Put("/posts/{id}", repository.UpdatePost)
	//router.Delete("/posts/{id}", repository.DeletePost)
	return router
}

func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r) // dispatch the request
	})
}
