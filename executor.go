package main

import "fmt"

type Executor interface {
    Submit(task func()) (int, error)
}

type InlineExecutor struct {}
type TaskItem struct {
    id int
    task func()
}
type ForegroundExecutor struct {
    uniqueId int
    tasklist []TaskItem
}

func (ie InlineExecutor) Submit(task func()) (int, error) {
    task()
    return 1, nil
}

func (fe ForegroundExecutor) Submit(task func()) (int, error) {
    fe.uniqueId++
    fe.tasklist = append(fe.tasklist, TaskItem { fe.uniqueId, task })
    return fe.uniqueId, nil
}

func (fe ForegroundExecutor) Execute() {
    for _, task := range fe.tasklist {
        task.task()
    }
}

func main() {
    fmt.Printf("hello, world\n")
    ie := InlineExecutor{}
    ie.Submit(func() {
        fmt.Printf("hello, inline\n")
    })

    fe := ForegroundExecutor{ 0 , make([]TaskItem, 1) } 
    fe.Submit(func() {
        fmt.Printf("hello, foreground %v\n", fe.uniqueId)
    })
    fe.Submit(func() {
        fmt.Printf("hello, foreground %v\n", fe.uniqueId)
    })

    fe.Execute()
}

