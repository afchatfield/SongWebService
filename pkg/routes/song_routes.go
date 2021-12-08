package routes

import (
	"github.com/afchatfield/SongWebService/pkg/controllers"
	"github.com/gorilla/mux"
)

// Route declaration
func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Index).Methods("GET")
	r.HandleFunc("/", controllers.Search).Methods("POST")
	r.HandleFunc("/index", controllers.Index).Methods("GET")
	r.HandleFunc("/index", controllers.Search).Methods("POST")
	r.HandleFunc("/show", controllers.Show)
	r.HandleFunc("/new", controllers.New).Methods("GET")
	r.HandleFunc("/insert", controllers.Insert).Methods("POST")
	r.HandleFunc("/edit", controllers.Edit).Methods("GET")
	r.HandleFunc("/update", controllers.Update).Methods("POST")
	r.HandleFunc("/delete", controllers.Delete)
	return r
}
