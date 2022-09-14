package routes

import (
	"github/shayansh001/bookstoremanagement/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book{bookID}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book{bookID}", controllers.UpdateBook).Methods("DELETE")
}
