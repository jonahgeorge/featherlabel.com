package lib

import (
	"database/sql"
	"fmt"
)

type sqlString sql.NullString

func (s sqlString) MarshalJSON() ([]byte, error) {
	if s.Valid == false {
		return []byte("\"\""), nil
	} else {
		return []byte(fmt.Sprintf("\"%s\"", s.String)), nil
	}
}

type sqlInt64 sql.NullInt64

func (s sqlInt64) MarshalJSON() ([]byte, error) {
	if s.Valid == false {
		return []byte("\"\""), nil
	} else {
		return []byte(fmt.Sprintf("\"%d\"", s.Int64)), nil
	}
}
