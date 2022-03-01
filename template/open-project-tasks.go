package template

import (
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
)

func getOpenTasks(header string, issues []jira.Issue) (template string) {
	template = fmt.Sprintf("============ %v ===========\n", header)
	for _, rx := range issues {
		//template += fmt.Sprintf("------------------------------\n")
		template += fmt.Sprintf(`Summary: %v
Last Updated: %v
Status; %v\n`, rx.Fields.Summary, rx.Fields.Updated, rx.Fields.Status.Name)
		template += fmt.Sprintf("Description:\n#+BEGIN_SRC quote\n%v\n#+END_SRC\n", rx.Fields.Description)
		template += fmt.Sprintf("------------------------------\n")
	}

	return
}

func GetOpenTasksInProject(proj string, issues []jira.Issue) (template string) {
	header := fmt.Sprintf("Open Tasks for Project: %v", proj)
	template = getOpenTasks(header, issues)
	return
}
