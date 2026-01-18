package model

type DateForBD struct {
	Year  string
	Month string
	Day   string
}

var Date = map[int64]*DateForBD{}
