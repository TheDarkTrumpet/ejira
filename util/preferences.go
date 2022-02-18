package util

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type GHVars struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Server   string `json:"server"`
}

func (vars GHVars) LoadPreferences(creds string) error {
	user, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	varsFile := fmt.Sprintf("%s/.creds/%s", user, creds)

	//var vars GHVars
	contents, err := ioutil.ReadFile(varsFile)
	err = json.Unmarshal(contents, &vars)
	if err != nil {
		return err
	}
	return nil
}
