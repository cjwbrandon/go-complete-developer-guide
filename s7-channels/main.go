package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Receiving messages
	// 	for i := 0; i < len(links); i++ {
	// 		fmt.Println(<-c)
	// 	}

	// Looping infinitely
	// for {
	// 	go checkLink(<-c, c)
	// }

	// Loop with each message
	for l := range c {
		// time.Sleep(time.Second * 5) // Do not put in main routine -> blocking/throttling
		// Use function literal -> rmb to pass by value
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}
