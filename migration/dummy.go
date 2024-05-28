package main

import (
	"csvdump/employee"
	"encoding/csv"
	"flag"
	"fmt"
	"github.com/go-faker/faker/v4"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	defaultDSN := "admin:admin@tcp(127.0.0.1:3307)/hr"
	dsn := flag.String("dsn", defaultDSN, "database connection string")
	flag.Parse()

	rand.New(rand.NewSource(time.Now().UnixNano()))

	file, err := os.Create("dummy_data.csv")
	if err != nil {
		panic(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for i := 1; i <= 1000000; i++ {
		emp := randomEmp()
		if err := emp.Save(*dsn); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err.Error())
		}
		_ = writer.Write(emp.ToSlice())
	}

	fmt.Println("CSV file successfully generated.")
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
