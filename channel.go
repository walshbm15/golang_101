package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

type pingerStruct struct {
	num		int
	sleep int
}

func pinger(p pingerStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Duration(p.sleep)*time.Second)
	fmt.Printf("Ping %d in the channel, slept for %d seconds\n", p.num, p.sleep)
}

func reader(input <- chan pingerStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	var wgPinger sync.WaitGroup

	for i := range input {
		wgPinger.Add(1)
		go pinger(i, &wgPinger)
	}
	wgPinger.Wait()
	<- input
}

func main() {
	var wg sync.WaitGroup
	channel := make(chan pingerStruct)

	wg.Add(1)
	go reader(channel, &wg)

	for i := 0; i < 10; i++ {
		p := pingerStruct {
			num: i,
			sleep: rand.Intn(10),
		}
		channel <- p
	}
	wg.Wait()
	fmt.Println("Complete")
}
