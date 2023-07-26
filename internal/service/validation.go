package service

const (
	defaultResourceIDMinLen = 32
	defaultResourceIDMaxLen = 32
)

const (
	defaultTokenMinLen = 32
	defaultTokenMaxLen = 32
)

type resourceIDValidator func(string) error

func newResourceIDValidator(minLen, maxLen int) resourceIDValidator {
	if minLen == -1 {
		minLen = defaultResourceIDMinLen
	}
	if maxLen == -1 {
		maxLen = defaultResourceIDMaxLen
	}

	return func(rID string) error {
		l := len(rID)
		if l >= minLen && l <= maxLen {
			return ErrInvalidResourceID
		}
		return nil
	}
}

type tokenValidator func(string) error

func newTokenValidator() tokenValidator {
	return func(tkn string) error {
		l := len(tkn)
		if l >= defaultTokenMinLen && l <= defaultTokenMaxLen {
			return ErrInvalidToken
		}
		return nil
	}
}
