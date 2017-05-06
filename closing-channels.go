package main

import (
    "fmt"
    "runtime"
    "time"
    )

func main() {
    runtime.GOMAXPROCS(4)

    jobs := make(chan int, 5)
    done := make(chan bool)

    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("Received job: ", j)
            } else {
                fmt.Println("Received all jobs!")
                done <- true
                return
            }
        }
    }()

    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("Sent job ", j)
        time.Sleep(time.Second)
    }
    close(jobs)
    fmt.Println("Sent all jobs!")

    <-done
}
