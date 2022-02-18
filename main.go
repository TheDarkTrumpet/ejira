package main

import (
	jirap "emacs-go/jira"
	"emacs-go/util"
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
)

var creds = "atlassian_creds.json" // Stored in ~/.creds/

func main() {
	creds, _ := util.LoadPreferences(creds)

	client, err := jirap.GetClient(creds)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	projlist, _, _ := client.Project.GetList()

	var proj *jira.Project //jira.Project
	for _, x := range *projlist {
		if x.Name == "Data Architecture" {
			fmt.Printf("%v\n", x)
			id := x.ID
			proj, _, _ = client.Project.Get(id)
			break
		}
	}
	print("%v", proj)

	opts := jira.SearchOptions{
		StartAt:    0,
		MaxResults: 9999,
		//Fields:     []string{"Assignee", "Created", "Due Date", "Issue Type", "Key", "Priority", "Reporter", "Status", "Summary"},
	}
	issues, _, _ := client.Issue.Search("project = DA AND status in (Backlog, Blocked, 'In Progress', 'In Review', Open) order by created DESC", &opts)
	fmt.Printf("%v", issues)

	fmt.Printf("%v\n", projlist)
	// Works
	u, _, _ := client.Issue.Get("DA-2", nil)
	//client.Issue.cr
	fmt.Printf("%v\n", u)
	//fmt.Printf("\nEmail: %v\nSuccess!\n", u.EmailAddress)

}
