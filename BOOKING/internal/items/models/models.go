package models

type Status string

const (
	CONFIRMED  = Status("CONFIRMED")
	REJECTED   = Status("REJECTED")
	INPROGRESS = Status("INPROGRESS")
)
