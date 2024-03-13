package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"time"
)

func inner(name string) func() {
	return func() {
		fmt.Println(name, time.Now())
	}
}

func f1() {
	fmt.Println("f1...", time.Now())
}

type Job struct {
	Name string
	Age  int
}

func (this Job) Run() {
	fmt.Printf("%s %s", this.Name, time.Now())
}

func main() {
	Cron := cron.New(cron.WithSeconds())
	Cron.AddFunc("*/2 * * * * *", inner("xiaoxin"))
	Cron.AddJob("*/3 * * * * *\n", Job{"xiaoxin", 18})
	Cron.Start()
	select {}
}
