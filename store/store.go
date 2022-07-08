package store

import (
	"io/ioutil"
	"time"

	"github.com/dowerx/goSchedule/ec"
	"gopkg.in/yaml.v2"
)

type TaskConfig struct {
	Type    string
	Command []string
	Token   string
}

type Task struct {
	Type string
	From time.Time
	To   time.Time
}

func LoadTaskConfigs(path string) []TaskConfig {
	data, err := ioutil.ReadFile(path)
	ec.Check(err)
	tasks := make([]TaskConfig, 0)
	err = yaml.Unmarshal(data, &tasks)
	ec.Check(err)
	return tasks
}

func LoadTasks(path string) []Task {
	data, err := ioutil.ReadFile(path)
	ec.Check(err)
	tasks := make([]Task, 0)
	err = yaml.Unmarshal(data, &tasks)
	ec.Check(err)
	return tasks
}

func SaveTasks(path string, tasks []Task) {
	data, err := yaml.Marshal(tasks)
	ec.Check(err)
	err = ioutil.WriteFile(path, data, 0)
	ec.Check(err)
}
