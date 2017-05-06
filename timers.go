package main

import (
    "fmt"
    "time"
)

func main() {
    timer1 := time.NewTimer(time.Second * 2)

    <-timer1.C
    fmt.Println("Timer 1 Expired")
    timer2 := time.NewTimer(time.Second)
    go func() {
        <-timer2.C
        fmt.Println("Timer 2 Expired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 Stopped")
    }
}
