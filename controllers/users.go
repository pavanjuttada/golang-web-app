package controllers

import (
	"text/template"
	"net/http"
	"fmt"
    "github.com/gorilla/mux"
	"../models"
)

var templates = template.Must(template.ParseFiles("templates/edit.html", "templates/view.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *models.Pagedata) {
	// p.Datasql = template.HTML()

    err := templates.ExecuteTemplate(w, tmpl+".html", p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
    
    vars := mux.Vars(r)
    fmt.Println(vars["id"]);
	   
    page_data := models.EditUser(vars["id"])
    w.Header().Set("Content-Type", "text/html")
    renderTemplate(w, "edit", page_data)
}

func ViewHandler(w http.ResponseWriter, r *http.Request) {
	page_data := models.GetAllUsers()
	w.Header().Set("Content-Type", "text/html")
	renderTemplate(w, "view", page_data)
}
func UpdateHandler(w http.ResponseWriter, r *http.Request){
    id := r.FormValue("id")
    name := r.FormValue("name")
    email := r.FormValue("email")
    address := r.FormValue("address")
	
	if models.UpdateUser(id, name, email, address){
		http.Redirect(w, r, "/view/", http.StatusFound)
	}else{
		fmt.Println("Failed to update the user")
	}
}
func SaveHandler(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")
    email := r.FormValue("email")
    address := r.FormValue("address")
    
	if models.CreateUser(name, email, address){
		http.Redirect(w, r, "/view/", http.StatusFound)
	}else{
		fmt.Println("Failed to create user")
	}
}
func DeleteHandler(w http.ResponseWriter, r *http.Request){

    vars := mux.Vars(r)
    fmt.Println(vars["id"]);

    if models.DeleteUser(vars["id"]) {
		http.Redirect(w, r, "/view/", http.StatusFound)
	}else{
		fmt.Println("Failed to delte user")
	}
}
