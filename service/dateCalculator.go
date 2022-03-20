package service

import (
	"nextPayday/model"
	"time"
)

func GetCurrentDate() time.Time {
	now := time.Now()

	return now
}

func handleWeekendPayday(payDay time.Time) time.Time {
	if payDay.Weekday().String() == "Saturday" {
		payDay = payDay.AddDate(0, 0, -1)
	}

	if payDay.Weekday().String() == "Sunday" {
		payDay = payDay.AddDate(0, 0, -2)
	}

	return payDay
}

func GetNextPayday(now time.Time, payDay int) time.Time {
	var payMonth time.Month

	if now.Day() < payDay {
		payMonth = now.Month()
	} else {
		nextMonth := now.AddDate(0, 1, 0)
		payMonth = nextMonth.Month()
	}

	payYear := now.Year()

	if now.Month().String() == "December" && payMonth.String() == "January" {
		payYear = now.Year() + 1
	}

	nextPayday := Date(payYear, int(payMonth), payDay) // WTF time.Month is int, but I still have to explicitly convert it?
	nextPayday = handleWeekendPayday(nextPayday)

	if nextPayday.Before(now) {
		// Weird case here, in case the payday is in a weekend, causing it to be payed earlier, but earlier than the current day we are on now
		nextPayday = Date(nextPayday.Year(), int(nextPayday.Month())+1, payDay)
		nextPayday = handleWeekendPayday(nextPayday)
	}

	return nextPayday
}

func GetNumberOfDaysUntilPayday(now time.Time, payDay time.Time) int {
	elapsed := payDay.Sub(now)

	return int(elapsed.Hours()) / 24
}

func GetSalariesThisYear(now time.Time, payday int) []model.PayDay {
	currentYear := now.Year()

	payDaysThisYear := make([]model.PayDay, 0, 12)

	auxToday := GetCurrentDate()
	for auxToday.Year() < currentYear+1 && auxToday.Month().String() != "December" {
		nextPayDay := GetNextPayday(auxToday, payday)
		days := GetNumberOfDaysUntilPayday(now, nextPayDay)

		payDay := model.NewPayDay(nextPayDay, days)
		payDaysThisYear = append(payDaysThisYear, payDay)

		auxToday = nextPayDay.AddDate(0, 0, 3)
	}

	return payDaysThisYear
}

func Date(year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
}
