package faker

import (
	"strconv"
	"strings"
)

// JiraIssueKey returns random Jira issue key
func JiraIssueKey() string {
	projectFormat := strings.Repeat("*", IntBetween(2, 5))
	project := UpperAsciify(projectFormat)
	number := IntBetween(1, 999)
	return project + "-" + strconv.Itoa(number)
}
