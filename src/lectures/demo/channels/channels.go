package main

import (
	"fmt"
	"time"
)

type ControlMsg int

const (
	DoExit = iota
	ExitOK
)

type Job struct {
	data int
}

type Result struct {
	result int
	job    Job
}

func Doubler(jobs <-chan Job, results chan<- Result, control chan ControlMsg) {
	for {
		select {
		case msg := <-control:
			switch msg {
			case DoExit:
				fmt.Println("Exit goroutine")
				control <- ExitOK
			default:
				panic("unhandled control message")
			}
		case job := <-jobs:
			results <- Result{result: job.data * 2, job: job}
		}
	}
}

func main() {
	jobs := make(chan Job, 50)
	results := make(chan Result, 50)
	control := make(chan ControlMsg)

	go Doubler(jobs, results, control)

	for i := 0; i < 30; i++ {
		jobs <- Job{i}
	}

	for {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-time.After(500 * time.Millisecond):
			fmt.Println("timed out")
			control <- DoExit
			<-control
			fmt.Println("program exit")
			return
		}
	}
}
