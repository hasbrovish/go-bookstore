package routes

import (
	"github.com/gorilla/mux"
	"github.com/hasbrovish/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).Methods("DELETE")
}
