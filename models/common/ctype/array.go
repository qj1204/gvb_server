package ctype

import (
	"database/sql/driver"
	"strings"
)

type Array []string

func (a *Array) Scan(value interface{}) error {
	v, _ := value.([]byte)
	if string(v) == "" {
		*a = []string{}
		return nil
	}
	*a = strings.Split(string(v), "\n")
	return nil
}

func (a Array) Value() (driver.Value, error) {
	return strings.Join(a, "\n"), nil
}
