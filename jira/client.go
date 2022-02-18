package jira

import (
	"emacs-go/util"
	"gopkg.in/andygrunwald/go-jira.v1"
)

type EJIRA struct {
	Creds  util.GHVars
	Client *jira.Client
}

func (ejira *EJIRA) GetClient() (err error) {
	if ejira.Client != nil {
		return nil // Only one instance of client, don't reconnect
	}

	tp := jira.BasicAuthTransport{
		Username: ejira.Creds.Username,
		Password: ejira.Creds.Password,
	}

	client, err := jira.NewClient(tp.Client(), ejira.Creds.Server)
	ejira.Client = client
	return
}
