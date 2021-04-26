package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func getter(url string, c chan string) {
	res, err := myClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	c <- string(data)
}

func main() {
	var c chan string = make(chan string)

	url := "https://jsonplaceholder.typicode.com/posts/"

	for i := 1; i < 101; i++ {
		urlput := url + strconv.Itoa(i)
		go getter(urlput, c)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}
}
