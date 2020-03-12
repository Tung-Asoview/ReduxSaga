package router

import (
	"Go-Chi/controller"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
	"time"
)

var router *chi.Mux

func Post_router(controller controller.Post) {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	router.Get("/posts", controller.AllPosts)
	//router.Get("/posts/{id}", controller.DetailPost)
	//router.Post("/posts", controller.CreatePost)
	//router.Put("/posts/{id}", controller.UpdatePost)
	//router.Delete("/posts/{id}", controller.DeletePost)
}

func Logger() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(time.Now(), r.Method, r.URL)
		router.ServeHTTP(w, r) // dispatch the request
	})
}