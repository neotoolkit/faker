package faker_test

import (
	"testing"

	"github.com/neotoolkit/faker"
)

func TestJira_Key(t *testing.T) {
	k := faker.NewFaker().Jira().Key()

	if len(k) == 0 {
		t.Error("jira issue key is empty")
	}
}
