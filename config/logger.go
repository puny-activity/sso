package config

type Logger struct {
	level string
}

func (c Logger) Level() string {
	return c.level
}
