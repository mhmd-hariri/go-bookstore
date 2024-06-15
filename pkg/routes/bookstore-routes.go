package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mhmd-hariri/go-bookstore/pkg/controllers"
	"github.com/mhmd-hariri/go-bookstore/pkg/handlers"
	"github.com/mhmd-hariri/go-bookstore/pkg/middleware"
)

var RegisterBookStoreRouters = func(router *mux.Router) {
	router.HandleFunc("/login", handlers.Login).Methods("POST")
	router.HandleFunc("/logout", handlers.Logout).Methods("POST")
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.Handle("/book", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateBook))).Methods("POST")
	router.Handle("/book", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetBooks))).Methods("GET")
	router.Handle("/book/{bookId}", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetBookById))).Methods("GET")
	router.Handle("/book/{bookId}", middleware.AuthMiddleware(http.HandlerFunc(controllers.UpdateBook))).Methods("PUT")
	router.Handle("/book/{bookId}", middleware.AuthMiddleware(http.HandlerFunc(controllers.DeleteBook))).Methods("DELETE")
}
