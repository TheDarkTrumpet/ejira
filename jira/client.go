package jira

import (
	"emacs-go/util"
	"gopkg.in/andygrunwald/go-jira.v1"
)

type EJIRA struct {
	Creds  util.GHVars
	Client *jira.Client
}

func (ejira EJIRA) GetClient() (err error) {
	tp := jira.BasicAuthTransport{
		Username: ejira.Creds.Username,
		Password: ejira.Creds.Password,
	}

	ejira.Client, err = jira.NewClient(tp.Client(), ejira.Creds.Server)
	return
}
