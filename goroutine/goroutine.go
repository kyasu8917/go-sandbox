package main

import (
    "fmt"
    "time"
)

func f() {
    time.Sleep(time.Second)
    fmt.Println("process done")
}


func main() {

    fmt.Println("start")

    for i:= 0; i < 10; i++ {
        go f()
    }

    fmt.Println("processing...")

    time.Sleep(10 * time.Second)
    fmt.Println("end")

}