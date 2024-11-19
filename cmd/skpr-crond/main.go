package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/skpr/crond/pkg/config"
	"github.com/alecthomas/kingpin/v2"
	"gopkg.in/yaml.v2"
)

var (
	cliConfig = kingpin.Flag("config", "Path to the crond config file.").Short('c').Default(config.DefaultFilePath).String()
)

func main() {
	kingpin.Parse()

	var tasks config.Tasks

	fmt.Println("Loading configuration")

	yamlFile, err := ioutil.ReadFile(*cliConfig)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &tasks)
	if err != nil {
		panic(err)
	}

	for name, task := range tasks {
		fmt.Printf("Found task: name='%s' schedule='%s' command='%s' args='%s'\n", name, task.Schedule, task.Command, strings.Join(task.Args, " "))
	}

	fmt.Println("Starting application")

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
