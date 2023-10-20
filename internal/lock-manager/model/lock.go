package model

import "github.com/gofrs/uuid/v5"

// Lock is a pair of resource and token to release it.
type Lock struct {
	// ResourceID is a unique identifier of a Lock.
	resourceID string

	// Token is a secret you can use to unlock the resource.
	token uuid.UUID
}

// ReinstateLock validates `rID` & `t` and returns lock or error if any.
func ReinstateLock(rID, t string) (*Lock, error) {
	tkn, err := uuid.FromString(t)
	if err != nil {
		return nil, err
	}
	return &Lock{
		resourceID: rID,
		token:      tkn,
	}, nil
}

// NewLock validates `rID` and returns created Lock or error if any.
func NewLock(rID string) (*Lock, error) {
	t, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return &Lock{
		resourceID: rID,
		token:      t,
	}, nil
}

func (l Lock) ResourceID() string {
	return l.resourceID
}

func (l Lock) Token() string {
	return l.token.String()
}
