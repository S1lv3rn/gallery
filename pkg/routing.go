package pkg

import (
    "net/http"
)




// ROUTER
func Start() { 
    fs := http.FileServer(http.Dir("assets/"))
    

    r := http.NewServeMux()
    r.HandleFunc("/", getHome)
    r.HandleFunc("/project/{id}", getProject)
    r.Handle("/static/", http.StripPrefix("/static/", fs))
    
    http.ListenAndServe(":8080", r)
}