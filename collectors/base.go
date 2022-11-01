package collectors

import (
	"database/sql"
	"fmt"
)

type baseCollector struct {
	db *sql.DB
}

func newBaseCollector(db *sql.DB) *baseCollector {
	return &baseCollector{db}
}

func (c *baseCollector) status(name string) float64 {
	row := c.db.QueryRow("show global status where variable_name = ?", name)

	var (
		vname string
		value float64
	)

	if err := row.Scan(vname, value); err == nil {
		return value
	}
	return 0
}

func (c *baseCollector) variables(name string) float64 {
	row := c.db.QueryRow("show variables where Variable_name = ?", name)
	var (
		vname string
		value float64
	)

	if err := row.Scan(vname, value); err == nil {
		fmt.Println(value)
		return value
	} else {
		fmt.Println(err)
		return 1
	}

	return 0
}
