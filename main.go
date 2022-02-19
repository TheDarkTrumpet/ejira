package main

import (
	jirap "emacs-go/jira"
	"emacs-go/util"
	"flag"
	"fmt"
	"log"
	"os"
)

//var creds_file = "atlassian_creds.json" // Stored in ~/.creds/
var credsFile *string
var operation *string

func init() {
	operation = flag.String("operation", "", "Operation to Perform")
	credsFile = flag.String("creds", "atlassian_creds.json", "Creds file to load (default atlassian_creds.json)")
}

func main() {
	flag.Parse()

	if len(*operation) == 0 || len(*credsFile) == 0 {
		fmt.Println("Usage: ejira -operation <operation_to_do>")
		flag.PrintDefaults()
		os.Exit(1)
	}

	creds, err := util.LoadPreferences(*credsFile)
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
