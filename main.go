package main

import (
	"github.com/ismdeep/justoj-monitor/config"
	"github.com/ismdeep/justoj-monitor/task"
	"time"
)

func pause() {
	for {
		time.Sleep(1 * time.Second)
	}
}

func main() {
	config.LoadConfig()
	task.StartPendingMonitor()
	pause()
}
