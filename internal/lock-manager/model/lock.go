package model

type Lock struct {
	// ResourceID is a unique identifier of a Lock.
	ResourceID string
	// Token is a secret you can use to unlock the resource.
	Token string
}
