package redisconn

// Config stores all configurable parameters for redis connection.
type Config struct {
	URI      string `env:"URI"   env-required:"true"`
	Username string `env:"UNAME"                     env-default:""`
	Password string `env:"PWORD"                     env-default:""`
	DB       uint   `env:"DB"                        env-default:"0"`
}
