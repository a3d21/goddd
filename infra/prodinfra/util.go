package prodinfra

import (
	mydr "github.com/go-sql-driver/mysql"
	"github.com/mattn/go-sqlite3"
)

// Duplicated 判断GormErr是否唯一索引冲突
func Duplicated(err error) bool {
	if err != nil {
		if driverError, ok := err.(*mydr.MySQLError); ok {
			return driverError.Number == 1062
		}

		if sqlerr, ok := err.(sqlite3.Error); ok {
			return sqlerr.ExtendedCode == 2067
		}
	}

	return false
}
