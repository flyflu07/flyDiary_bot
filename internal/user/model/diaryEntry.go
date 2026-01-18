package model

import "time"

type DiaryEntry struct {
	Id      string
	Message string
	Time    time.Time
}
