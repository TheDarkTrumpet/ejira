package template

import (
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
	"regexp"
)

func GetOrgDetails(issue *jira.Issue) string {
	re := regexp.MustCompile("[[:^ascii:]]")
	var comments string
	assignee := "Unassigned"
	reportedBy := "Unknown"

	if issue.Fields.Comments != nil {
		comments = fmt.Sprintf("\n#+begin_quote\n=========== COMMENTS ===========\n")
		for x := 0; x < len(issue.Fields.Comments.Comments); x++ {
			c := issue.Fields.Comments.Comments[x]
			comments += fmt.Sprintf(`------
Date: %v
Author: %v
Comment: %v
-----`, c.Created, c.Author.DisplayName, re.ReplaceAllLiteralString(c.Body, ""))
		}
		comments += fmt.Sprintf("\n========= END COMMENTS =========\n#+end_quote\n")
	}

	if issue.Fields.Assignee != nil {
		assignee = issue.Fields.Assignee.DisplayName
	}

	if issue.Fields.Reporter != nil {
		reportedBy = issue.Fields.Reporter.DisplayName
	}

	description := re.ReplaceAllLiteralString(issue.Fields.Description, "")

	text := fmt.Sprintf(`Issue: %v : %v
Description: %v
Assigned To: %v
Reported By: %v
Status: %v
Comments:%v`, issue.Key, issue.Fields.Summary,
		description,
		assignee, reportedBy,
		issue.Fields.Status.Name,
		comments)

	return text
}
