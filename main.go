package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
	"io/ioutil"
	"log"
	"os"
)

var creds = "atlassian_creds.json" // Stored in ~/.creds/

func main() {
	creds, _ := loadVars()

	tp := jira.BasicAuthTransport{
		Username: creds.Username,
		Password: creds.Password,
	}

	client, err := jira.NewClient(tp.Client(), creds.Server)

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
	for x := 0; x < len(*projlist); x++ {

	}

	fmt.Printf("%v\n", projlist)
	// Works
	u, _, _ := client.Issue.Get("DA-2", nil)
	//client.Issue.cr
	fmt.Printf("%v\n", u)
	//fmt.Printf("\nEmail: %v\nSuccess!\n", u.EmailAddress)

}

type GHVars struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Server   string `json:"server"`
}

func loadVars() (GHVars, error) {
	user, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	varsFile := fmt.Sprintf("%s/.creds/%s", user, creds)

	var vars GHVars
	contents, err := ioutil.ReadFile(varsFile)
	err = json.Unmarshal(contents, &vars)
	if err != nil {
		return vars, err
	}
	return vars, err
}
