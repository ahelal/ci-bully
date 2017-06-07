package config

// Config configuration options
type Config struct {
	Token   string      `yaml:"token"`
	Actions ActionSlice `yaml:"actions"`
	Repos   []string    `yaml:"repos"`
}

type Action struct {
	Day     int    `yaml:"day"`
	Action  string `yaml:"action"`
	Message string `yaml:"message"`
}

type ActionSlice []Action
