package main

import (
	"time"
)

// GetYMD return with string: YYYYMMDD
func GetYMD() string {
	t := time.Now().Local()
	s := t.Format("20060102")
	return s
}
