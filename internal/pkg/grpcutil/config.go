package grpcutil

type Config struct {
	Port           string `yaml:"port"`
	WithReflection bool   `yaml:"withReflection"`
}
