package controllers

import (
	"fmt"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Unable to parse from submission", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "<p>Email: %s</p>", r.PostForm.Get("email"))
	fmt.Fprintf(w, "<p>Password: %s</p>", r.PostForm.Get("password"))
}
