package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

var emps []employee

func main() {
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		expensiveCall()
		fmt.Printf("druation is: %+vs\n", time.Now().Sub(start).Seconds())
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type employee struct {
	Name     string
	Age      int
	Title    *string
	Birthday time.Time
	Content  []byte
}

func expensiveCall() {
	for i := 0; i < 100; i++ {
		emp := getEmployees()
		emps = append(emps, emp...)
	}
}

func getEmployees() []employee {
	var f []employee
	for i := 0; i < 100; i++ {
		title := "ceo"
		bir := time.Now()
		data, err := ioutil.ReadFile("./../../redis.pdf")
		if err != nil {
			log.Fatal(err)
		}
		t := employee{
			Name:     "adam",
			Age:      23,
			Title:    &title,
			Birthday: bir,
			Content:  data,
		}
		f = append(f, t)
	}
	return f
}
