package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func update(wg *sync.WaitGroup, host, version string) {
	n := rand.Intn(100) + 50
	spent := time.Duration(n) * time.Millisecond

	time.Sleep(spent)
	log.Printf("%s updated to %s in %.2fs", host, version, spent.Seconds())

	wg.Done()
}

func updateAll(version string, hosts <-chan string) {
	var wg sync.WaitGroup

	for host := range hosts {
		wg.Add(1)
		go update(&wg, host, version)
	}

	wg.Wait()
}

func scanVMs(hosts chan<- string) {
	for i := 0; i < 5; i++ {
		host := fmt.Sprintf("srv%d", i+1)
		hosts <- host
	}
	close(hosts)
}

func patchVMs(version string) string {
	hosts := make(chan string)

	go scanVMs(hosts)

	updateAll(version, hosts)

	return "all servers updated"
}
