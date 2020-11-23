package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/twinj/uuid"
)

type Task struct {
	Id uuid.UUID
	IsCompleted bool // have a random function to mark the IsCompleted after a random period  
	Status string //completed, failed, timeout
	Time time.Time // when was the task created
	TaskData string // random string indicating task
}

func enqueue(queue[] Task, task Task) []Task {
	queue = append(queue, task); // Simply append the task to enqueue.
	fmt.Printf("Enqueued: %v", task);
	return queue
}

func dequeue(queue []Task) ([]Task) {
	task := queue[0]; // The first task in the queue is the one to be dequeued.
	fmt.Printf("Dequeued: %v", task)
	return queue[1:]; // Slice off the element once it is dequeued.
}

func main() {
	var queue[] Task; // Make a queue of Task.
	//Create 9 tasks
	for i:=0; i<9 ; i++ {
		taskData := fmt.Sprintf("Creating email task %d", i)
		queue = append(queue, Task{
			//generating a UUID for each task
			Id:          uuid.NewV1(),
			// Setting the iscompleted flag of every task as false
			IsCompleted: false,             // have a random function to mark the IsCompleted after a random period  
			// by default setting all task as failed
			Status:      "failed",         //completed, failed, timeout
			Time:        time.Now(),        // when was the task created
			TaskData:    taskData, // random string indicating task
		},
		)
	}
	fmt.Printf("Inital queue %v", queue)
	go pollTask(queue)
	// Un comment the below sleep to see the logs of go routines
	//time.Sleep(10 * time.Minute)
}

func pollTask(queue []Task){
	for len(queue) > 0 {
		status := processTaskAndSendStatus(queue[0])
		//check the difference in time here we are keeping the timeout period as 900 seconds, which means if the
		// task is more than  900 seconds in a queue then we mark that task as "timeout"
		diff := subtractTime(queue[0].Time, time.Now())
		if diff <= 900 {
			if status == "completed" {
				queue[0].IsCompleted = true
				queue[0].Status = "completed"
				queue = dequeue(queue)
			} else if status == "failed" {
				queue = enqueue(queue, queue[0])
			}
		} else {
			queue[0].Status = "timeout"
			queue = dequeue(queue)
		}
		fmt.Printf("********************Printing the queue***************** %v", queue)
		//the poll happens on every 10 second interval
		time.Sleep(10 * time.Second)
	}
}

func processTaskAndSendStatus(task Task) string {
	//creates only two random numbers either 0 or 1
	n := rand.Intn(2)
	if n == 0 {
		return "completed"
	}
	return "failed"
}

func subtractTime(time1,time2 time.Time) float64{
	diff := time2.Sub(time1).Seconds()
	return diff
}