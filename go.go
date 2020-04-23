package main

import (
	"database/sql"
	"fmt"

	log "github.com/sirupsen/logrus"
)

type config struct {
	URL  string `json:"url"`
	Rank int    `json:"rank"`
}

func (c *config) url() string {
	db, err := sql.Open("sqlite3", "/var/lib/db/example.db")
	defer db.Close()

	if err != nil {
		log.Fatal("Failed")
	}
	return c.URL
}

func function1(v1 int, c chan int) int {
	fmt.Println("function 1")

	v2 := <-c

	if v1 < 10 && v2 > 20 {
		fmt.Println("Line")
	}
	return 1
}

func function2(c chan int) {
	c <- 30
}

func main() {
	myChan := make(chan int)
	go function2(myChan)
	function1(1, myChan)
}
