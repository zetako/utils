# utils

Small Go utility helpers.

Current API:

- `ContextValue[T]`: typed getter for `context.Context` values.

## Install

```bash
go get github.com/zetako/utils
```

## Example

```go
package main

import (
	"context"
	"fmt"

	"github.com/zetako/utils"
)

type userIDKey struct{}

func main() {
	ctx := context.WithValue(context.Background(), userIDKey{}, 42)

	userID, ok := utils.ContextValue[int](ctx, userIDKey{})
	fmt.Println(userID, ok)
}
```