package service

const (
	minResourceIDLength = 32
	maxResourceIDLength = 32
)

const (
	minTokenLength = 32
	maxTokenLength = 32
)

func isValidResourceID(rID string) bool {
	l := len(rID)
	return l >= minResourceIDLength && l <= maxResourceIDLength
}

func isValidToken(tkn string) bool {
	l := len(tkn)
	return l >= minTokenLength && l <= maxTokenLength
}
