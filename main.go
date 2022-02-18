package main

import (
	jirap "emacs-go/jira"
	"emacs-go/util"
	"log"
	"os"
)

var creds_file = "atlassian_creds.json" // Stored in ~/.creds/

func main() {
	creds, err := util.LoadPreferences(creds_file)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ejira := jirap.EJIRA{Creds: creds}

	proj, _ := ejira.GetProjectByName("Data Architecture")
	print("%v", proj)

	/*
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
	*/
}
