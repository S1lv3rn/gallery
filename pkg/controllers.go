package pkg

// CONTROLLERS
/*
Controllers:
 - gets request
 - extracts post/get variables and vaildates
 - requests model data
 - put model data in view

*/

import (
    "net/http"
    "html/template"
    "strconv"
)

var funcMap = template.FuncMap{
    "formatDate": formatDate,
}


func getProject(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.Atoi(r.PathValue("id"))
    if checkToDisplay(err, w) {return}

    data, err := getProjectData(id, w)
    if checkToDisplay(err, w) {return}


    tmpl := template.Must(template.New("project.html").Funcs(funcMap).ParseFiles(
        "views/project.html",
        "views/common.html",
    ))
    tmpl.Execute(w, data)
}

func getHome(w http.ResponseWriter, r *http.Request) {
    data := getHomeData()


    tmpl := template.Must(template.New("gallery.html").Funcs(funcMap).ParseFiles(
        "views/gallery.html",
        "views/common.html",
    ))
    tmpl.Execute(w, data)
}
