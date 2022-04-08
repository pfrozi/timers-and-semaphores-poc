package main

import (
	"fmt"
	"time"
)

const TimerSec = 4
const MaxPerTimerSec = 2
const PoolSize = 10

var ps [10]int
var ch chan int
var semaphore chan int

var timer1 *time.Timer

func main() {

	timer1 = time.NewTimer(TimerSec * time.Second)

	semaphore = make(chan int, PoolSize-1)
	ch = make(chan int)
	ps = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

	go runWorker()

	// wait to start
	fmt.Println("Waiting the start . . . ")
	time.Sleep(3 * time.Second)

	fmt.Println("Starting the process!")
	fmt.Println(". . .")

	for i, p := range ps {
		fmt.Printf("--> Sending p(%d) to the channel . . .\n", p)
		ch <- p
		// wait (P)
		semaphore <- 1

		if (i+1)%MaxPerTimerSec == 0 {
			fmt.Println("Wating after ", i+1)
			// waiting the timer's end
			<-timer1.C
			fmt.Println("Wating Complete!")
			// restarting the timer
			timer1.Reset(TimerSec * time.Second)
		}
	}

	// Waiting all processes
	time.Sleep(20 * time.Second)
}

func runWorker() {
	for p := range ch {
		go process(p)
	}
}

func process(p int) {
	fmt.Printf("--> --> Processing p(%d) . . .\n", p)
	time.Sleep(1 * time.Second)
	fmt.Printf("--> --> p(%d) processed !\n", p)
	// signal / release (V)
	<-semaphore
}
