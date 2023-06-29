package model

type Lock struct {
	// ResourceID is a unique identifier of a Lock.
	ResourceID string
	// SecretKey is a secret key you can use to unlock the resource.
	SecretKey string
}
