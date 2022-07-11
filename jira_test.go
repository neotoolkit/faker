package faker_test

import (
	"testing"

	"github.com/neotoolkit/faker"
)

func TestJiraIssueKey(t *testing.T) {
	k := faker.JiraIssueKey()

	if len(k) == 0 {
		t.Error("jira issue key is empty")
	}
}
