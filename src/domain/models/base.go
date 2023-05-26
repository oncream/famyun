package models

import "time"

type DBTime struct {
	Created time.Time `json:"-"`
	Updated time.Time `json:"-"`
	Deleted time.Time `json:"-"`
}
