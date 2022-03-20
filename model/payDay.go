package model

import (
	"fmt"
	"time"
)

type PayDay struct {
	Date      string
	DaysUntil int
}

func NewPayDay(date time.Time, daysUntil int) PayDay {
	self := PayDay{}

	self.Date = fmt.Sprintf("%s %d, %d", date.Month(), date.Day(), date.Year())
	self.DaysUntil = daysUntil

	return self
}

func (self *PayDay) String() string {
	return fmt.Sprintf("Pay day is in %d day(s), on %s",
		self.DaysUntil,
		self.Date,
	)
}
