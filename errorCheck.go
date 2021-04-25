package exercise

import (
	"database/sql"
	"errors"
	"fmt"
)

type DaoError struct {
	sql string
	err error
}

func (de *DaoError) Error() string {
	return fmt.Sprintf("exercise.Daoerror : sql:[%s]\r\n", de.sql)
}
func (de *DaoError) Unwrap() error {
	return de.err
}
func Query(sqlstr string) error {
	er := sql.ErrNoRows
	return &DaoError{err: er, sql: sqlstr}
}

func main() {
	err := Query("select * form user ")
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("no rows")
	}

}
