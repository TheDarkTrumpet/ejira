package jira

import (
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
)

func (ejira *EJIRA) GetIssuesByProject(project *jira.Project) (issues []jira.Issue, err error) {
	ejira.GetClient()
	opts := jira.SearchOptions{
		StartAt:    0,
		MaxResults: 9999,
	}
	jql := fmt.Sprintf("project = %s AND status in (Backlog, Blocked, 'In Progress', 'In Review', Open) order by created DESC", project.Key)
	issues, _, err = ejira.Client.Issue.Search(jql, &opts)

	return
}
