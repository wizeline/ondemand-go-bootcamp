package customtype

import "fmt"

var _ fmt.Stringer = DriverDB("")

const (
	UndefinedDriverDB DriverDB = "undefined"
	CsvDriverDB       DriverDB = "csv"
)

type DriverDB string

func (d DriverDB) String() string {
	return string(d)
}

func NewDriverDB(driver string) DriverDB {
	switch driver {
	case CsvDriverDB.String():
		return CsvDriverDB
	default:
		return UndefinedDriverDB
	}
}
