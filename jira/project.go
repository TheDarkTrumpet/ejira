package jira

import (
	"gopkg.in/andygrunwald/go-jira.v1"
)

func (ejira EJIRA) getProjects() (jira.ProjectList, error) {
	projlist, _, err := ejira.Client.Project.GetList()
	return *projlist, err
}

func (ejira EJIRA) GetProjectByName(name string) (*jira.Project, error) {
	ejira.GetClient()

	projlist, err := ejira.getProjects()

	if err != nil {
		return nil, err
	}

	var proj *jira.Project
	for _, x := range projlist {
		if x.Name == name {
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
