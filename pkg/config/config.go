package config

// DefaultFilePath loaded when the application starts.
const DefaultFilePath = "/etc/skpr/data/cron.yml"

// Tasks which are executed during the lifecycle of this application.
type Tasks map[string]Task

// Task which is executed during the lifecycle of this application.
type Task struct {
	Schedule string   `json:"schedule" yaml:"schedule"`
	Command  string   `json:"command" yaml:"command"`
	Args     []string `json:"args" yaml:"args"`
}
