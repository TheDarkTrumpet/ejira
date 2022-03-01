package template

import (
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
)

func GetOpenTasksInProject(issues []jira.Issue) (template string) {
	template = fmt.Sprintf("============================ Open Tasks ===========================\n")
	for _, rx := range issues {
		template += fmt.Sprintf("-------------------------------------------------------------------\n")
		template += fmt.Sprintf(`Summary: %v
Last Updated: %v
Status; %v\n`, rx.Fields.Summary, rx.Fields.Updated, rx.Fields.Status.Name)
		template += fmt.Sprintf("Description:\n#+BEGIN_SRC quote\n%v\n#+END_SRC\n", rx.Fields.Description)
		template += fmt.Sprintf("-------------------------------------------------------------------\n")
	}

	return
}
