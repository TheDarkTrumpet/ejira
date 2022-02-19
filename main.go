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
var value *string

func init() {
	operation = flag.String("operation", "", "Operation to Perform")
	value = flag.String("value", "", "The value tied to the operation, depends on the context")
	credsFile = flag.String("creds", "atlassian_creds.json", "Creds file to load (default atlassian_creds.json)")
}

var allowableOperations = map[string]string{
	"OpenTasks":        "Retrieve all open tasks assigned to the currently logged in user (value flag can be blank/null)",
	"OpenProjectTasks": "Retrieve all open tasks in a project (defined by value flag)",
	"OrgJiraDetails":   "Retrieve a formatted entry that can be inserted into org-mode, by task id (defined by value flag)",
}

func main() {
	flag.Parse()

	if len(*operation) == 0 || len(*credsFile) == 0 {
		fmt.Println("Usage: ejira -operation <operation_to_do>")
		flag.PrintDefaults()
		fmt.Println("---------------------------------------")
		fmt.Println("---- Allowable Operations and Help ----")
		fmt.Println("---------------------------------------")
		for k, v := range allowableOperations {
			fmt.Printf("%s   :   %v\n", k, v)
		}
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
}

func OpenTasks() {

}

func OpenProjectTasks() {

}

func OrgJiraDetails() {

}
