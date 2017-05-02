package main

import "fmt"

type Executor interface {
    Submit(task func()) (int, error)
}

type InlineExecutor struct {}

func (ie InlineExecutor) Submit(task func()) (int, error) {
    task()
    return 1, nil
}

func main() {
    fmt.Printf("hello, world\n")
    ie := InlineExecutor{}
    ie.Submit(func() {
        fmt.Printf("hello, inline\n")
    })
}

