package jira

import (
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
	"io/ioutil"
	"strings"
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

	comment.Body = fmt.Sprintf("{code:text}%s{code}", string(fcontent))

	comment.Author = *me
	_, _, err = ejira.Client.Issue.AddComment(id, &comment)
	return
}

func getDescriptionSummary(file string) (summary string, description string) {
	lines := strings.Split(file, "\n")

	description = "{code:text}\n"
	for i, line := range lines {
		if i == 0 {
			summary = strings.Replace(line, "* ", "", 1)
		} else {
			description += fmt.Sprintf("%s\n", line)
		}
	}
	description += "{code}\n"
	return
}

func (ejira *EJIRA) AddIssue(proj string, file string) (result string, err error) {
	ejira.GetClient()

	fcontent, err := ioutil.ReadFile(file)

	if err != nil {
		return
	}

	summary, description := getDescriptionSummary(string(fcontent))

	me, err := ejira.GetCurrentUser()
	if err != nil {
		return
	}

	project, err := ejira.GetProjectByName(proj)
	if err != nil {
		return
	}

	issue := jira.Issue{
		Fields: &jira.IssueFields{
			Description: description,
			Summary:     summary,
			Project: jira.Project{
				Key: project.Key,
			},
			Type: jira.IssueType{
				Name: "Task",
			},
		},
	}

	basicIssue, _, err := ejira.Client.Issue.Create(&issue)
	ejira.Client.Issue.UpdateAssignee(basicIssue.ID, me)

	result = basicIssue.Key

	return
}
