package main

import "fmt"

func main() {
    messages := make(chan string)
    signals := make(chan bool)

    select{
    case msg := <-messages:
        fmt.Println("Received Message", msg)
    default:
        fmt.Println("No message received")
    }

    msg := "Hi"
    select {
    case messages <- msg:
        fmt.Println("Message sent:", msg)
    default:
        fmt.Println("No message sent.")
    }

    select {
    case msg := <- messages:
        fmt.Println("Received message: ", msg)
    case sig := <-signals:
        fmt.Println("Signal received: ", sig)
    default:
        fmt.Println("No activity")
    }
}
