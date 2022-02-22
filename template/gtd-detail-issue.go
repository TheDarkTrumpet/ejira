package template

import (
	"fmt"
	"gopkg.in/andygrunwald/go-jira.v1"
)

func GetOrgDetails(issue *jira.Issue) string {
	var comments string
	assignee := "Unassigned"

	if issue.Fields.Comments != nil {
		comments = fmt.Sprintf("\n#+begin_quote\n=========== COMMENTS ===========\n")
		for x := 0; x < len(issue.Fields.Comments.Comments); x++ {
			c := issue.Fields.Comments.Comments[x]
			comments += fmt.Sprintf(`------
Date: %v
Author: %v
Comment: %v
-----`, c.Created, c.Author.DisplayName, c.Body)
		}
		comments += fmt.Sprintf("\n========= END COMMENTS =========\n#+end_quote\n")
	}

	if issue.Fields.Assignee != nil {
		assignee = issue.Fields.Assignee.DisplayName
	}

	text := fmt.Sprintf(`Issue: [[%v][%v]]
Description: %v
Assigned To: %v
Status: %v
Comments: %v`, "Link", issue.Fields.Summary,
		issue.Fields.Description, assignee,
		issue.Fields.Status.Name, comments)

	return text
}
