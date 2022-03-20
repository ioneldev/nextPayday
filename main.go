package main

import (
	"fmt"
	"log"
	"net/http"
	"nextPayday/controller"
	"regexp"
)

var validPath = regexp.MustCompile("^/(next_salary|until_eoy_salary)")

func main() {
	fmt.Println("Start")

	http.HandleFunc("/next_salary", makeGetHandler(controller.NextSalaryHandler))
	http.HandleFunc("/until_eoy_salary", makeGetHandler(controller.UntilEoySalaryHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func makeGetHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)

		if r.Method != "GET" {
			http.NotFound(w, r)
			return
		}
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}
