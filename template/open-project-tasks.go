package template

import (
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
)

func GetOpenTasksInProject(issues []jira.Issue) (template string) {
	for _, rx := range issues {
		fmt.Println(rx)
	}

	return
}
