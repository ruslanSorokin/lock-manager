package test

func Must[T any](t *T, u error) *T {
	if u != nil {
		panic(u)
	}
	return t
}
