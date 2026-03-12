# utils

[![Go Reference](https://pkg.go.dev/badge/github.com/zetako/utils.svg)](https://pkg.go.dev/github.com/zetako/utils)
[![Go Report Card](https://goreportcard.com/badge/github.com/zetako/utils)](https://goreportcard.com/report/github.com/zetako/utils)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub Tag](https://img.shields.io/github/v/tag/zetako/utils?sort=semver)](https://github.com/zetako/utils/tags)
[![Commit Messages](https://img.shields.io/badge/commit%20message-gitmoji%20unicode-informational)](./.github/workflows/commit-message.yml)

Small, focused Go utility helpers.

The repository starts with typed helpers around `context.Context` and is intended to grow as a small collection of narrowly scoped utilities instead of a grab bag package.

## Install

```bash
go get github.com/zetako/utils@latest
```

## Package docs

- Go package docs: https://pkg.go.dev/github.com/zetako/utils
- Module path: `github.com/zetako/utils`

## API

### ContextValue

`ContextValue[T]` reads a value from `context.Context`, attempts a type assertion to `T`, and returns the typed value plus a success flag.

Use typed, unexported context keys to avoid collisions:

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
	if !ok {
		panic("missing user id")
	}

	fmt.Println(userID)
}
```

If the key is missing or the stored value is not assignable to `T`, `ContextValue` returns the zero value of `T` and `false`.

## Commit message policy

This repository requires gitmoji commit messages in unicode mode.

Valid examples:

- `✨ add ContextValue helper`
- `📝 expand README and package docs`
- `✅ add tests for missing context values`

The policy is enforced in two places:

- locally through the repository commit hook in `.githooks/commit-msg`
- remotely through GitHub Actions in `.github/workflows/commit-message.yml`

To enable the local hook after cloning:

```bash
git config core.hooksPath .githooks
```

## License

MIT. See [LICENSE](./LICENSE).