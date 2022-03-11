package jira

import (
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
	"io/ioutil"
)

func (ejira *EJIRA) GetIssuebyID(id string) (issue *jira.Issue, err error) {
	ejira.GetClient()

	issue, _, err = ejira.Client.Issue.Get(id, nil)
	return
}

func (ejira *EJIRA) GetMyIssues() (issues []jira.Issue, err error) {
	ejira.GetClient()
	opts := jira.SearchOptions{
		StartAt:    0,
		MaxResults: 9999,
	}

	jql := "assignee in (currentUser()) AND status in (Backlog, Blocked, 'In Progress', 'In Review', Open) order by created DESC"
	issues, _, err = ejira.Client.Issue.Search(jql, &opts)

	return
}

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

func (ejira *EJIRA) PutCommentToIssue(id string, file string) (err error) {
	ejira.GetClient()

	fcontent, err := ioutil.ReadFile(file)

	if err != nil {
		return
	}

	me, err := ejira.GetCurrentUser()
	if err != nil {
		return
	}

	var comment jira.Comment

	comment.Body = string(fcontent)
	comment.Author = *me
	_, _, err = ejira.Client.Issue.AddComment(id, &comment)
	return
}
