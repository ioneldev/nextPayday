package service

import (
	"fmt"
	"nextPayday/model"
	"reflect"
	"testing"
	"time"
)

func TestGetNextPayday(t *testing.T) {
	var tests = []struct {
		payDay   int
		from     time.Time
		expected time.Time
	}{
		{18, Date(2022, 3, 17), Date(2022, 3, 18)},
		{17, Date(2022, 3, 19), Date(2022, 4, 15)},
		{1, Date(2022, 5, 30), Date(2022, 6, 1)},
		{27, Date(2022, 7, 30), Date(2022, 8, 26)},
		{15, Date(2022, 3, 15), Date(2022, 4, 15)},
		{20, Date(2022, 3, 19), Date(2022, 4, 20)},
		{29, Date(2022, 2, 19), Date(2022, 3, 1)},
	}
	for _, data := range tests {
		testName := fmt.Sprintf("Pay day:%d,Current date: %v", data.payDay, data.from)
		t.Run(testName, func(t *testing.T) {
			ans := GetNextPayday(data.from, data.payDay)
			if ans != data.expected {
				t.Errorf("Actual %v, Expected %v", ans, data.expected)
			}
		})
	}
}

func TestGetNumberOfDaysUntilPayday(t *testing.T) {
	var tests = []struct {
		payDate  time.Time
		from     time.Time
		expected int
	}{
		{Date(2022, 3, 17), Date(2022, 3, 18), 1},
		{Date(2022, 3, 19), Date(2022, 4, 17), 29},
		{Date(2022, 5, 30), Date(2022, 6, 1), 2},
		{Date(2022, 7, 30), Date(2022, 8, 27), 28},
		{Date(2022, 3, 15), Date(2022, 4, 15), 31},
	}
	for _, data := range tests {
		testName := fmt.Sprintf("Pay day:%v,Current date: %v", data.payDate, data.from)
		t.Run(testName, func(t *testing.T) {
			ans := GetNumberOfDaysUntilPayday(data.payDate, data.from)
			if ans != data.expected {
				t.Errorf("Actual %v, Expected %d", ans, data.expected)
			}
		})
	}
}

func TestGetSalariesThisYear(t *testing.T) {
	var tests = []struct {
		payDay   int
		from     time.Time
		expected []model.PayDay
	}{
		{18, Date(2022, 3, 17), []model.PayDay{
			{"April 18, 2022", 32},
			{"May 18, 2022", 62},
			{"June 17, 2022", 92},
			{"July 18, 2022", 123},
			{"August 18, 2022", 154},
			{"September 16, 2022", 183},
			{"October 18, 2022", 215},
			{"November 18, 2022", 246},
			{"December 16, 2022", 274},
		}},
	}
	for _, data := range tests {
		testName := fmt.Sprintf("Pay day:%d,Current date: %v", data.payDay, data.from)
		t.Run(testName, func(t *testing.T) {
			ans := GetSalariesThisYear(data.from, data.payDay)
			if !reflect.DeepEqual(ans, data.expected) {
				t.Errorf("Actual %v, Expected %v", ans, data.expected)
			}
		})
	}
}
