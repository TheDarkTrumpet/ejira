package jira

import (
	"emacs-go/util"
	"gopkg.in/andygrunwald/go-jira.v1"
)

func GetClient(creds util.GHVars) (client *jira.Client, err error) {
	tp := jira.BasicAuthTransport{
		Username: creds.Username,
		Password: creds.Password,
	}

	client, err = jira.NewClient(tp.Client(), creds.Server)
	return
}
