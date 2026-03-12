package utils

import "context"

// ContextValue retrieves a value from ctx by key and attempts to return it as T.
//
// If the key does not exist or the stored value is not assignable to T, it
// returns the zero value of T and false.
func ContextValue[T any](ctx context.Context, key any) (T, bool) {
	value := ctx.Value(key)
	typedValue, ok := value.(T)
	if !ok {
		var zero T
		return zero, false
	}

	return typedValue, true
}
