package sql

import (
	"fmt"

	"github.com/xwb1989/sqlparser"
)

// Parse wraps the sqlparser library and gives us access to it
func Parse(input string) error {
	stmt, err := sqlparser.Parse(input)
	if err != nil {
		return err
	}

	// Otherwise do something with stmt
	switch stmt := stmt.(type) {
	case *sqlparser.Select:
		fmt.Printf("Select: %+v\n", stmt)
	case *sqlparser.Insert:
		fmt.Printf("Insert: %+v\n", stmt)
	}

	return err
}
