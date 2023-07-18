package service

const (
	defaultResourceIDMinLen = 32
	defaultResourceIDMaxLen = 32
)

const (
	defaultTokenMinLen = 32
	defaultTokenMaxLen = 32
)

type resourceIDValidator func(string) bool

func newResourceIDValidator(minLen, maxLen int) resourceIDValidator {
	if minLen == -1 {
		minLen = defaultResourceIDMinLen
	}
	if maxLen == -1 {
		maxLen = defaultResourceIDMaxLen
	}

	return func(rID string) bool {
		l := len(rID)
		return l >= minLen && l <= maxLen
	}
}

type tokenValidator func(string) bool

func newTokenValidator(minLen, maxLen int) tokenValidator {
	if minLen == -1 {
		minLen = defaultTokenMinLen
	}
	if maxLen == -1 {
		maxLen = defaultTokenMaxLen
	}

	return func(tkn string) bool {
		l := len(tkn)
		return l >= minLen && l <= maxLen
	}
}
