package main

import (
	"fmt"
	"time"
	"sync"
)

func worker(number int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d starting \n", number)

	time.Sleep(time.Second)

	fmt.Printf("Worker %d done\n", number)

}

func main() {

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()

	fmt.Println("All done")

}