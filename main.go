package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/csv"
	"flag"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

//go:embed help.txt
var helpMsg string

func main() {
	var (
		help     bool
		socket   string
		host     string
		port     int
		user     string
		password string
		database string
		table    string
	)
	flag.BoolVar(&help, "help", false, "-help to get help")
	flag.StringVar(&socket, "socket", "", "-socket socket")
	flag.StringVar(&host, "host", "127.0.0.1", "-host host")
	flag.IntVar(&port, "port", 3306, "-port port")
	flag.StringVar(&user, "user", "", "-user username")
	flag.StringVar(&password, "password", "", "-password password")
	flag.StringVar(&database, "database", "", "-database database name")
	flag.StringVar(&table, "table", "", "-table table name")
	flag.Parse()
	if help {
		fmt.Println(helpMsg)
		os.Exit(0)
	}
	dsn, err := buildDSN(socket, host, port, user, password, database)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// get all employees
	data, err := getData(dsn, table)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	if err := csvPrint(data); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
	}
}

func csvPrint(data [][]string) error {
	csvWriter := csv.NewWriter(os.Stdout)
	err := csvWriter.WriteAll(data)
	if err != nil {
		return err
	}
	csvWriter.Flush()
	if err := csvWriter.Error(); err != nil {
		return fmt.Errorf("failed to write CSV: %s", err.Error())
	}
	return nil
}

//go:embed get_data_query.txt
var GetDataQuery string

func getData(dsn string, table string) ([][]string, error) {
	var out [][]string
	query := fmt.Sprintf(GetDataQuery, table)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return out, fmt.Errorf("failed to open database: %s", err.Error())
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return out, fmt.Errorf("data get error: %s" + err.Error())
	}
	columns, errColumns := rows.Columns()
	out = append(out, columns)
	if errColumns != nil {
		return out, fmt.Errorf("columns scan error: %s", errColumns.Error())
	}
	for rows.Next() {
		colPointers := make([]interface{}, len(columns))
		pointerContainer := make([]string, len(columns))
		for i := range colPointers {
			colPointers[i] = &pointerContainer[i]
		}
		err = rows.Scan(colPointers...)
		if err != nil {
			return out, fmt.Errorf("failed to scan data: %s", err.Error())
		}
		out = append(out, pointerContainer)
	}
	if err := rows.Err(); err != nil {
		return out, fmt.Errorf("rows scan error: %s" + err.Error())
	}
	return out, nil
}

func buildDSN(socket string,
	host string,
	port int,
	user string,
	password string,
	database string) (string, error) {
	const (
		UNIXPROTOCOL = "unix"
		TCPPROTOCOL  = "tcp"
	)
	// return error if user empty
	if len(user) == 0 {
		return "", fmt.Errorf("user cannot be empty")
	}
	if len(socket) > 0 {
		return fmt.Sprintf("%s@%s(%s)/%s",
			user, UNIXPROTOCOL, socket, database), nil
	} else {
		return fmt.Sprintf("%s:%s@%s(%s:%d)/%s",
			user, password, TCPPROTOCOL, host, port, database), nil
	}
}
