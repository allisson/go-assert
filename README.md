# go-assert
[![Build Status](https://github.com/allisson/go-assert/workflows/Release/badge.svg)](https://github.com/allisson/go-assert/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/allisson/go-assert)](https://goreportcard.com/report/github.com/allisson/go-assert)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/allisson/go-assert)

Simple asserts to use with golang testing.

## Quickstart

```golang
// file test_test.go
package test

import (
    "testing"

    assert "github.com/allisson/go-assert"
)

func TestExamples(t *testing.T) {
    assert.Equal(t, 1, 1)
    assert.NotEqual(t, 1, 2)
    assert.Nil(t, nil)
    assert.NotNil(t, 1)
}
```

```bash
go test -v
=== RUN   TestExamples
--- PASS: TestExamples (0.00s)
PASS
```

```golang
// file test_test.go
package test

import (
    "testing"

    assert "github.com/allisson/go-assert"
)

func TestExamples(t *testing.T) {
    assert.Equal(t, 1, 2)
}
```

```bash
go test -v
=== RUN   TestExamples
--- FAIL: TestExamples (0.00s)
    test_test.go:10: assertion_type=Equal, expected_value=1, expected_type=int, current_value=2, current_type=int
FAIL
exit status 1
```
