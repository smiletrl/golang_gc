package main

import (
	_ "flag"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	_ "os"
	_ "runtime"
	_ "runtime/pprof"
	_ "sync"
	"time"
)

func main() {
	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

		start := time.Now()
		getLenth()

		end := time.Now()
		druation := end.Sub(start).Seconds()
		fmt.Printf("druation is: %+vs\n", druation)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type employee struct {
	Name     string
	Age      int
	Sex      *string
	Birthday time.Time
	Content  *[]byte
}

//go:noinline
func getLenth() {
	goroutine := 1000
	for i := 0; i < goroutine; i++ {
		emp := random()
		if len(*emp) > 10000 {
			fmt.Println(len(*emp))
		}
	}
}

//go:noinline
func random() *[]employee {
	var f []employee
	for i := 0; i < 10; i++ {
		sex := "man"
		bir := time.Now()
		dat, err := ioutil.ReadFile("./redis.pdf")
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Printf("bir is: %+v\n", bir)
		t := employee{
			Name:     "rulin",
			Age:      23,
			Sex:      &sex,
			Birthday: bir,
			Content:  &dat,
		}
		f = append(f, t)
	}
	return &f
}
