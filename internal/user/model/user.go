package model

import "time"

type User struct {
	Password           string
	TimeOfStartSession time.Time
	TimeOfAutoLogOut   time.Duration
}

var InfoAbUser = map[int64]*User{}
