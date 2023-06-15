package model

import "time"

type Lock struct {
	// ResourceID is an identefier of a locked resource.
	ResourceID string
	// SecretKey is a secret key you can use to unlock.
	SecretKey  string
	// Lifetime is a duration during which this resource will be locked unless you unlock it first or renew it.
	Lifetime   time.Duration
}
