# go-assert
[![Build Status](https://travis-ci.org/allisson/go-assert.svg)](https://travis-ci.org/allisson/go-assert)
[![Go Report Card](https://goreportcard.com/badge/github.com/allisson/go-assert)](https://goreportcard.com/report/github.com/allisson/go-assert)
[![Documentation](https://godoc.org/github.com/allisson/go-assert?status.svg)](http://godoc.org/github.com/allisson/go-assert)

Simple asserts to use with golang testing.

## Quickstart

```golang
package main

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
