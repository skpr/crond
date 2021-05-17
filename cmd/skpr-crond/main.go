package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/skpr/crond/internal/config"
	"gopkg.in/alecthomas/kingpin.v2"
	"gopkg.in/yaml.v2"
)

var (
	cliConfig = kingpin.Flag("config", "Path to the crond config file.").Short('c').Default("/etc/skpr/crond/tasks.yml").String()
)

func main() {
	var tasks config.Tasks

	yamlFile, err := ioutil.ReadFile(*cliConfig)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &tasks)
	if err != nil {
		panic(err)
	}

	if err := run(tasks); err != nil {
		panic(err)
	}
}

func run(tasks config.Tasks) error {
	scheduler := gocron.NewScheduler(time.UTC)

	for _, task := range tasks {
		scheduler.Cron(task.Schedule).Do(newTask(task.Command, task.Args...))
	}

	scheduler.StartBlocking()

	return nil
}

func newTask(command string, args ...string) func() {
	return func() {
		cmd := exec.Command(command, args...)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
