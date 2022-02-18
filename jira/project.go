package jira

import (
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
)

func (ejira EJIRA) GetProjects() (projlist *jira.ProjectList, err error) {
	projlist, _, err = ejira.Client.Project.GetList()
	return
}

func (ejira EJIRA) GetProjectByName(name string) (*jira.Project, error) {
	projlist, err := ejira.GetProjects()

	if err != nil {
		return nil, err
	}
	
	var proj *jira.Project
	for _, x := range *projlist {
		if x.Name == name {
			fmt.Printf("%v\n", x)
			id := x.ID
			proj, _, err = ejira.Client.Project.Get(id)
			break
		}
	}

	if err != nil {
		return nil, err
	}

	return proj, nil
}
