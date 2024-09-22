package pkg

import (
    "fmt"
    "time"
    "net/http"
)



// COMMON STRUCTS

type Page struct {
    PageTitle string
    Projects []Project
}

type Project struct  {
    Id int `json:"id"`
    Title string `json:"title"`
    Created int64 `json:"created"`
    Updated int64 `json:"updated"`
    NoteTexts []string `json:"noteTexts"`
}

type HomePage struct {
    Page
}

type ProjectPage struct {
    Page
    Project
}

type ProjectsJson struct {
    SiteUpdated int64 `json:"lastUpdated"`
    Projects []Project `json:"projects"`
}


// COMMON FUNCTIONS

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func checkToDisplay(e error, w http.ResponseWriter) bool {
    if e != nil {
        fmt.Fprintf(w, e.Error())
    }
    return e != nil
}

func formatDate(epochTime int64) string {
    return time.Unix(epochTime, 0).Format("01/02/2006")
}