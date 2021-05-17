package config

type Tasks map[string]Task

type Task struct {
	Schedule string   `json:"schedule" yaml:"schedule"`
	Command  string   `json:"command" yaml:"command"`
	Args     []string `json:"args" yaml:"args"`
}
