package main

import (
    "fmt"
    "time"
    "sync/atomic"
    "runtime"
)

func main() {
    runtime.GOMAXPROCS(50)

    var ops uint64 = 0

    for i := 0; i < 50; i++ {
        go func() {
            for {
                atomic.AddUint64(&ops, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)

    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops: ", opsFinal)
}
