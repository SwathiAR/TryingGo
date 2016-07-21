package main

import (
	"fmt"


	"time"

)

var i int = 0;

func main() {
	ch1 := make(chan string, 10)
         ch1 <- "Let the game bigin"


	go pinger(ch1)



     time.Sleep(time.Duration(1000000))


}

func pinger(ch chan string) {

	x := <-ch

	fmt.Println(x)
	ch <- "ping"
	go ponger(ch)

}

func ponger(ch chan string) {
	x := <-ch


	fmt.Println(x, i)
	ch <- "pong"
	i++;
	go pinger(ch)

}

