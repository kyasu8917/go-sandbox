package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("start")
	messages := make(chan string)
	
	go func() { 
		time.Sleep(time.Second)
		messages <- "ping"
	}()

	fmt.Println("processing...")

	msg := <- messages
	fmt.Println("goroutine end")
	fmt.Println(msg)
}

// >>start
// >>processing...
// >>goroutine end
// >>ping

// channelのreceiverはsenderの処理が終了するまで処理をブロックする（＝待つ）