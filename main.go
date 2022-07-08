package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/dowerx/goSchedule/config"
	"github.com/dowerx/goSchedule/ec"
	"github.com/dowerx/goSchedule/store"
	"github.com/dowerx/goSchedule/web"
)

var tasks []store.Task
var taskConfigs []store.TaskConfig
var cfg config.Config

var cfgPath = "./config.yml"

func main() {
	if p := os.Getenv("GOSCHEDULECONFIG"); p != "" {
		cfgPath = p
	}
	cfg = config.LoadConfig(cfgPath)

	taskConfigs = store.LoadTaskConfigs(cfg.TaskConfigs)
	tasks = store.LoadTasks(cfg.Tasks)

	for _, tc := range taskConfigs {
		fmt.Printf("taskconfig: %s, %s, %s\n", tc.Type, tc.Command, tc.Token)
	}

	for _, t := range tasks {
		fmt.Printf("task: %s, %d\n", t.Type, t.From)
		go waitUntil(t)
	}

	web.Init(addTask, listTasks)
	web.Listen(cfg.Address)
}

func addTask(t string, df string, dt string, k string) {
	fmt.Printf("add: %s, %s, %s, %s\n", t, df, dt, k)
	if t != "" && df != "" && dt != "" && k != "" {
		for _, tc := range taskConfigs {
			if tc.Type == t && tc.Token == k {
				dateFrom, err := time.Parse(time.RFC3339, df)
				if err != nil {
					return
				}
				dateTo, err := time.Parse(time.RFC3339, dt)
				if err != nil {
					return
				}
				temp := store.Task{Type: t, From: dateFrom, To: dateTo}
				tasks = append(tasks, temp)
				store.SaveTasks(cfg.Tasks, tasks)
				go waitUntil(temp)
			}
		}
	}
}

func listTasks() []byte {
	data, err := json.Marshal(tasks)
	ec.Check(err)
	return data
}

func waitUntil(task store.Task) {
	untilFrom := time.Until(task.From)
	untilTo := time.Until(task.To)
	if untilFrom < untilTo && untilFrom > 0 {
		fmt.Printf("wait to start: %s\n", untilFrom)
		time.Sleep(untilFrom)
		for _, tc := range taskConfigs {
			if tc.Type == task.Type {
				fmt.Printf("run: %s: %s\n", tc.Type, tc.Command)
				cmd := exec.Command(tc.Command[0], tc.Command[1:]...)
				err := cmd.Start()
				ec.Check(err)
				untilTo = time.Until(task.To)
				fmt.Printf("wait to stop: %s\n", untilTo)
				time.Sleep(untilTo)
				fmt.Printf("stop: %s: %s\n", tc.Type, tc.Command)
				cmd.Process.Kill()
			}
		}
	}

	for i, t := range tasks {
		if t == task {
			tasks = remove(tasks, i)
			fmt.Println("remove")
			store.SaveTasks(cfg.Tasks, tasks)
		}
	}
}

func remove(slice []store.Task, s int) []store.Task {
	return append(slice[:s], slice[s+1:]...)
}
