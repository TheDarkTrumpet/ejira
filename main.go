package main

import (
	jirap "emacs-go/jira"
	"emacs-go/template"
	"emacs-go/util"
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
)

//var creds_file = "atlassian_creds.json" // Stored in ~/.creds/
var credsFile *string
var operation *string
var value *string
var valueFile *string

func init() {
	operation = flag.String("operation", "", "Operation to Perform")
	value = flag.String("value", "", "The value tied to the operation, depends on the context")
	valueFile = flag.String("vfile", "", "The path to the file that contains the value for operations (Mainly put operations)")
	credsFile = flag.String("creds", "atlassian_creds.json", "Creds file to load (default atlassian_creds.json)")
}

var allowableOperations = map[string]string{
	"OpenTasks":        "Retrieve all open tasks assigned to the currently logged in user (value flag can be blank/null)",
	"OpenProjectTasks": "Retrieve all open tasks in a project (defined by value flag)",
	"OrgJiraDetails":   "Retrieve a formatted entry that can be inserted into org-mode, by task id (defined by value flag)",
	"AddComment":       "Adds a comment to the task (defined by value flag)",
	"AddIssue":         "Creates a new issue based off the value (project), and temporary file as what to use as the body",
}

type T struct{}

func main() {
	flag.Parse()

	if len(*operation) == 0 || len(*credsFile) == 0 {
		PrintHelpAndExit()
	}

	creds, err := util.LoadPreferences(*credsFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	ejira := jirap.EJIRA{Creds: creds}

	if _, ok := allowableOperations[*operation]; !ok {
		log.Fatal(fmt.Sprintf("The operation, %v, was not found!  Please see help below:", *operation))
		PrintHelpAndExit()
	} else {
		var t T
		method := reflect.ValueOf(&t).MethodByName(*operation)
		mcall := make([]reflect.Value, method.Type().NumIn())
		mcall[0] = reflect.ValueOf(&ejira)
		mcall[1] = reflect.ValueOf(*value)
		output := method.Call(mcall)
		fmt.Printf("output: %v\n", output)
	}
}

func PrintHelpAndExit() {
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

func (t *T) OpenTasks(ejira *jirap.EJIRA, _ string) (temp string) {
	issues, err := ejira.GetMyIssues()
	if err != nil {
		log.Fatal(err)
		return
	}
	temp = template.GetOpenTasks(issues)
	return
}

func (t *T) OpenProjectTasks(ejira *jirap.EJIRA, val string) (temp string) {
	proj, err := ejira.GetProjectByName(val)
	if err != nil {
		log.Fatal(err)
		return
	}

	if proj == nil {
		log.Fatal(fmt.Sprintf("Unable to find a project by the name of %v", val))
		return
	}

	issues, err := ejira.GetIssuesByProject(proj)
	if err != nil {
		log.Fatal(err)
		return
	}

	temp = template.GetOpenTasksInProject(proj.Name, issues)
	return
}

// OrgJiraDetails takes an issue ID, and returns an org-compatible block for the details section.
func (t *T) OrgJiraDetails(ejira *jirap.EJIRA, val string) (orgDetails string) {
	issue, err := ejira.GetIssuebyID(val)

	if err != nil {
		log.Fatal(err)
		return
	}
	orgDetails = template.GetOrgDetails(issue)
	return
}

// AddComment takes an issue ID, and adds a comment to it
func (t *T) AddComment(ejira *jirap.EJIRA, val string) (err error) {
	err = ejira.PutCommentToIssue(val, *valueFile) // This kind of breaks my previous implementation pattern, may rethink this.

	if err != nil {
		log.Fatal(err)
	}
	return
}

// AddIssue takes a project ID (or name), and adds an issue to it
func (t *T) AddIssue(ejira *jirap.EJIRA, val string) (err error) {
	err = ejira.AddIssue(val, *valueFile) // This kind of breaks my previous implementation pattern, may rethink this.

	if err != nil {
		log.Fatal(err)
	}
	return
}
