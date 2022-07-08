package main

import (
	"fmt"

	"github.com/DowerX/record/store"
	"github.com/DowerX/record/web"
)

var tasks []store.Task
var taskConfigs []store.TaskConfig

func main() {
	taskConfigs = store.LoadTaskConfigs("./taskconfigs.yml")
	tasks = store.LoadTasks("./tasks.yml")

	for _, tc := range taskConfigs {
		fmt.Printf("taskconfig: %s, %s, %s\n", tc.Type, tc.Command, tc.Token)
	}

	for _, t := range tasks {
		fmt.Printf("task: %s, %d\n", t.Type, t.Date)
	}

	web.Init(addTask)
	web.Listen()
}

func addTask(t string, d string, k string) {
	fmt.Printf("add: %s, %s, %s", t, d, k)
}
