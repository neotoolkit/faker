package faker

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// IntBetween returns a int between a given minimum and maximum values
func IntBetween(min, max int) int {
	diff := max - min

	if diff == 0 {
		return min
	}

	bigInt, err := rand.Int(rand.Reader, big.NewInt(int64(diff)+1))
	if err != nil {
		return diff
	}

	n := bigInt.Int64()

	return int(n) + min
}

// Number returns a number between a given minimum and maximum values
func Number(min, max int) int {
	return IntBetween(min, max)
}

// RandomStringElement returns a random string element from a given list of strings
func RandomStringElement(s []string) string {
	i := IntBetween(0, len(s)-1)
	return s[i]
}

// Asciify returns string that replace all "*" characters with random ASCII values from a given string
func Asciify(in string) string {
	var out strings.Builder

	out.Grow(len(in))

	for i := range in {
		if in[i] == '*' {
			if Boolean() {
				out.WriteString(fmt.Sprintf("%c", IntBetween(65, 90)))
			} else {
				out.WriteString(fmt.Sprintf("%c", IntBetween(97, 122)))
			}
		} else {
			out.WriteByte(in[i])
		}
	}

	return out.String()
}

// LowerAsciify returns string that replace all "*" characters with random lower ASCII values from a given string
func LowerAsciify(in string) string {
	var out strings.Builder

	out.Grow(len(in))

	for i := range in {
		if in[i] == '*' {
			out.WriteString(fmt.Sprintf("%c", IntBetween(97, 122)))
		} else {
			out.WriteByte(in[i])
		}
	}

	return out.String()
}

// UpperAsciify returns string that replace all "*" characters with random upper ASCII values from a given string
func UpperAsciify(in string) string {
	var out strings.Builder

	out.Grow(len(in))

	for i := range in {
		if in[i] == '*' {
			out.WriteString(fmt.Sprintf("%c", IntBetween(65, 90)))
		} else {
			out.WriteByte(in[i])
		}
	}

	return out.String()
}

// Numerify returns a fake string that replace all "*" characters with numbers from 0 to 9 as string for Faker
func Numerify(in string) string {
	var out strings.Builder

	out.Grow(len(in))

	for i := range in {
		if in[i] == '*' {
			out.WriteString(strconv.Itoa(IntBetween(0, 9)))
		} else {
			out.WriteByte(in[i])
		}
	}

	return out.String()
}

// Faker returns random data as string by faker name
func Faker(name string) (interface{}, error) {
	switch strings.ToLower(name) {
	// Boolean
	case "boolean":
		return Boolean(), nil
	// Internet
	case "username":
		return Username(), nil
	case "gtld":
		return GTLD(), nil
	case "domain":
		return Domain(), nil
	case "email":
		return Email(), nil
	// Person
	case "firstname", "person.firstname":
		return FirstName(), nil
	case "lastname", "person.lastname":
		return LastName(), nil
	case "firstname male", "person.firstnamemale":
		return FirstNameMale(), nil
	case "firstname female", "person.firstnamefemale":
		return FirstNameFemale(), nil
	case "name", "person.name":
		return Name(), nil
	case "name male", "person.namemale":
		return NameMale(), nil
	case "name female", "person.namefemale":
		return NameFemale(), nil
	case "gender", "person.gender":
		return Gender(), nil
	case "gender male", "person.gendermale":
		return GenderMale(), nil
	case "gender female", "person.genderfemale":
		return GenderFemale(), nil
	// UUID
	case "uuid":
		return UUID(), nil
	}

	if strings.HasPrefix(strings.ToLower(name), "number(") {
		min, max, err := getNumberArgs(name)
		if err != nil {
			return "", err
		}
		return Number(min, max), nil
	}

	return nil, nil
}

func getNumberArgs(name string) (int, int, error) {
	args := strings.Trim(name[7:], ")")
	splitArgs := strings.Split(args, ",")
	min, err := strconv.Atoi(strings.TrimSpace(splitArgs[0]))
	if err != nil {
		return 0, 0, err
	}
	max, err := strconv.Atoi(strings.TrimSpace(splitArgs[1]))
	if err != nil {
		return 0, 0, err
	}
	return min, max, nil
}
