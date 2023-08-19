package redisutil

// Config stores all configurable parameters for redis connection.
type Config struct {
	URI      string
	Username string
	Password string
	DB       uint
}
