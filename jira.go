package faker

import (
	"strconv"
	"strings"
)

type Jira struct {
	Faker *Faker
}

// Key returns random Jira issue key
func (j Jira) Key() string {
	projectFormat := strings.Repeat("*", j.Faker.IntBetween(2, 5))
	project := j.Faker.UpperAsciify(projectFormat)
	number := j.Faker.IntBetween(1, 999)
	return project + "-" + strconv.Itoa(number)
}
