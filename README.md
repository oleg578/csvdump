### Simple dump utility

## Usage

```
go run . -port 3307 -user admin -password admin -database hr -table employee > dump.csv
```


- -h help
- -help to get help
- -socket string socket
- -host string host (default "127.0.0.1")
- -port int port (default 3306)
- -user string username
- -password string password
- -database string database name
- -table string table name

## Standard way to get csv file export with header

- see sql/csv_dump.sql