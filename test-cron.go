package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

func task() {
	fmt.Println("Task is being performed")
}

func main() {
	s := gocron.NewScheduler()
	s.Every(5).Seconds().Do(task)
	<-s.Start()
}
