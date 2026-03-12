package utils

import "context"

func ContextValue[T any](ctx context.Context, key any) (T, bool) {
	value := ctx.Value(key)
	typedValue, ok := value.(T)
	if !ok {
		var zero T
		return zero, false
	}

	return typedValue, true
}
