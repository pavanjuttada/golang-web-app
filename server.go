package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"./controllers"
)

func main() {

    r := mux.NewRouter()
    r.HandleFunc("/", controllers.ViewHandler)
	r.HandleFunc("/view/", controllers.ViewHandler)
	r.HandleFunc("/edit/{id:[0-9]+}", controllers.EditHandler)
	r.HandleFunc("/save/", controllers.SaveHandler)
    r.HandleFunc("/update/", controllers.UpdateHandler)
    r.HandleFunc("/delete/{id:[0-9]+}", controllers.DeleteHandler)
    r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
    http.ListenAndServe(":8080", r)
}