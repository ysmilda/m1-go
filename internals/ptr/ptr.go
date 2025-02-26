package ptr

func For[T any](v T) *T {
	return &v
}
