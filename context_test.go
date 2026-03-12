package utils

import (
	"context"
	"testing"
)

func TestContextValue(t *testing.T) {
	t.Run("returns typed value when present", func(t *testing.T) {
		type keyType string

		ctx := context.WithValue(context.Background(), keyType("user_id"), 42)

		value, ok := ContextValue[int](ctx, keyType("user_id"))
		if !ok {
			t.Fatal("expected value to be present")
		}

		if value != 42 {
			t.Fatalf("expected 42, got %d", value)
		}
	})

	t.Run("returns zero value when key is missing", func(t *testing.T) {
		type keyType string

		value, ok := ContextValue[string](context.Background(), keyType("missing"))
		if ok {
			t.Fatal("expected missing value")
		}

		if value != "" {
			t.Fatalf("expected empty string, got %q", value)
		}
	})

	t.Run("returns zero value when type does not match", func(t *testing.T) {
		type keyType string

		ctx := context.WithValue(context.Background(), keyType("count"), 42)

		value, ok := ContextValue[string](ctx, keyType("count"))
		if ok {
			t.Fatal("expected type assertion to fail")
		}

		if value != "" {
			t.Fatalf("expected empty string, got %q", value)
		}
	})
}
