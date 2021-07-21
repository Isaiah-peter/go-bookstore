package routes

import (
	"github.com/Isaiah-peter/go-bookstore/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).methods("POST")
	router.HandleFunc("/book/", controllers.GetAllBooks).methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.GetBookById).methods("GET")
	router.HandleFunc("/book/{bookid}", controllers.UpdateBook).methods("PUT")
	router.HandleFunc("/book/{bookid}", controllers.DeleteBook).methods("DELETE")

}
