package service

type Config struct {
	ResourceID struct {
		MaxLen int `yaml:"maxLength" env-default:"-1"`
		MinLen int `yaml:"minLength" env-default:"-1"`
	} `yaml:"resourceID"`

	Token struct {
		MaxLen int `yaml:"maxLength" env-default:"-1"`
		MinLen int `yaml:"minLength" env-default:"-1"`
	} `yaml:"token"`
}
