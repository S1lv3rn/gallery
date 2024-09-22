package pkg

import (
    "encoding/json"
    "errors"
    "os"
)

// DATABASE MODELS


func getProjectsJson() []Project {
    projectItem := ProjectsJson{}
    jsonByt, err := os.ReadFile("data/projects.json")
    check(err)

    json.Unmarshal(jsonByt, &projectItem)
    //check(err)

    return projectItem.Projects
}

func getProjectById(id int) (Project, error) {
    projectList := getProjectsJson()
    project := Project{}
    err := errors.New("Project id not valid")

    for _, p := range projectList {
        if id == p.Id {
            err = nil
            project = p
            break
        }
    }
    return project, err
}
