package employee

import (
	"context"
	"database/sql"
	_ "embed"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Employee struct {
	ID        string `json:"id" faker:"uuid_hyphenated"`
	FirstName string `json:"first_name" faker:"first_name"`
	LastName  string `json:"last_name" faker:"last_name"`
	Email     string `json:"email" faker:"email"`
	Phone     string `json:"phone_number" faker:"phone_number"`
}

func (emp *Employee) ToSlice() []string {
	return []string{
		emp.ID,
		emp.FirstName,
		emp.LastName,
		emp.Email,
		emp.Phone,
	}
}

//go:embed save.sql
var SaveSql string

func (emp *Employee) Save(dsn string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return fmt.Errorf("failed to open database: %s", err.Error())
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	_, err = db.ExecContext(ctx, SaveSql,
		emp.ID,
		emp.FirstName,
		emp.LastName,
		emp.Email,
		emp.Phone,
		emp.FirstName,
		emp.LastName,
		emp.Email,
		emp.Phone)
	if err != nil {
		return fmt.Errorf("emp save error: %s" + err.Error())
	}
	return nil
}