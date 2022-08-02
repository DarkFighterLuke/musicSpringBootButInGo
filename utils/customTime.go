package utils

import (
	"fmt"
	"strings"
	"time"
)

const ctLayout = "02-01-2006"

var nilTime = (time.Time{}).UnixNano()

type CustomeTime struct {
	DataUscita time.Time
}

func (ct *CustomeTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.DataUscita = time.Time{}
		return
	}

	ct.DataUscita, err = time.Parse(ctLayout, s)
	return
}

func (ct *CustomeTime) MarshalJSON() ([]byte, error) {
	if ct.DataUscita.UnixNano() == nilTime {
		return []byte("null"), nil
	}

	return []byte(fmt.Sprintf("\"%s\"", ct.DataUscita.Format(ctLayout))), nil
}
