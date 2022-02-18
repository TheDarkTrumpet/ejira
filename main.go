package main

import (
	jirap "emacs-go/jira"
	"emacs-go/util"
	"fmt"
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
	fmt.Printf("%v\n", proj.Name)

	issues, _ := ejira.GetIssuesByProject(proj)

	fmt.Printf("%v\n", issues)

	i, _ := ejira.GetIssuebyID("DA-2")
	fmt.Printf("%v\n", i.Fields.Assignee)
	/*

		fmt.Printf("%v", issues)

		fmt.Printf("%v\n", projlist)
		// Works
		u, _, _ := client.Issue.Get("DA-2", nil)
		//client.Issue.cr
		fmt.Printf("%v\n", u)
		//fmt.Printf("\nEmail: %v\nSuccess!\n", u.EmailAddress)
	*/
}
