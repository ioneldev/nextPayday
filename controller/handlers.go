package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"nextPayday/model"
	"nextPayday/service"
	"strconv"
)

func NextSalaryHandler(w http.ResponseWriter, r *http.Request) {
	param, ok := r.URL.Query()["payday"]

	if !ok {
		log.Println("Url Param 'payday' is missing")
		w.WriteHeader(400)
		return
	}

	payday, err := strconv.Atoi(param[0])
	if err != nil {
		log.Println("Something wrong with atoi conversion", err)
		w.WriteHeader(500)
		return
	}

	log.Println("Url Param 'payday' is ", payday)

	nextSalary := GetNextSalary(payday)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(nextSalary)
	if err != nil {
		log.Println("Something wrong with json encode", err)
		w.WriteHeader(500)
		return
	}
}

func UntilEoySalaryHandler(w http.ResponseWriter, r *http.Request) {
	param, ok := r.URL.Query()["payday"]

	if !ok {
		log.Println("Url Param 'payday' is missing")
		w.WriteHeader(400)
		return
	}

	payday, err := strconv.Atoi(param[0])
	if err != nil {
		log.Println("Something wrong with atoi conversion", err)
		w.WriteHeader(500)
		return
	}

	log.Println("Url Param 'payday' is ", payday)

	today := service.GetCurrentDate()
	payDaysThisYear := service.GetSalariesThisYear(today, payday)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(payDaysThisYear)

	if err != nil {
		log.Println("Something wrong with json encode", err)
		w.WriteHeader(500)
		return
	}
}

func GetNextSalary(payday int) model.PayDay {
	today := service.GetCurrentDate()
	nextPayDay := service.GetNextPayday(today, payday)

	days := service.GetNumberOfDaysUntilPayday(today, nextPayDay)

	return model.NewPayDay(nextPayDay, days)
}
