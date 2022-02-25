package template

import (
	jirap "emacs-go/jira"
	"fmt"
)

func GetOpenTasksInProject(ejira *jirap.EJIRA, project string) (template string, err error) {
	proj, err := ejira.GetProjectByName(project)
	if err != nil {
		return
	}

	issues, err := ejira.GetIssuesByProject(proj)
	if err != nil {
		return
	}

	for _, rx := range issues {
		fmt.Println(rx)
	}

	return
}
