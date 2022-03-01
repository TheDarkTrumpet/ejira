package template

import (
	"gopkg.in/andygrunwald/go-jira.v1"
)

func GetOpenTasks(issues []jira.Issue) (template string) {
	header := "My Open Tasks"
	template = getOpenTasks(header, issues) // see open-project-tasks.go for this function
	return
}
