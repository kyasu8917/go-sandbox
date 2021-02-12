package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("start")

	channel := make(chan int)

	ids := []int {1,2,3,4,5,6,7,8,9,10}

	for _, id := range ids {
		go func (id int) {
			fmt.Printf("findById start id = %d\n", id)
		
			time.Sleep(time.Second)
		
			fmt.Printf("findById end id = %d\n", id)

			channel <- id * 10
		}(id)
	}		

	var results []int
	for i := 0; i < len(ids); i++ {
		result := <-channel
		results = append(results, result)
	}

	for _, result := range results {
		fmt.Printf("result = %d\n", result) 
	}
	
}