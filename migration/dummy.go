package main

import (
	"csvdump/employee"
	"flag"
	"fmt"
	"github.com/go-faker/faker/v4"
	"os"
	"strings"
)

func main() {
	defaultDSN := "admin:admin@tcp(127.0.0.1:3307)/hr"
	dsn := flag.String("dsn", defaultDSN, "database connection string")
	flag.Parse()

	for i := 1; i <= 1000000; i++ {
		emp := randomEmp()
		err := emp.Save(*dsn)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		}
	}
}

func randomEmp() employee.Employee {
	firstName := faker.FirstName()
	lastName := faker.LastName()
	return employee.Employee{
		ID:        0,
		FirstName: firstName,
		LastName:  lastName,
		Email: strings.Join(
			[]string{strings.Join([]string{
				strings.ToLower(firstName), strings.ToLower(lastName)}, "."), "corp.com"}, "@"),

		Phone: faker.Phonenumber(),
	}
}
