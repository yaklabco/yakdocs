//go:build stave

package main

import (
	"github.com/yaklabco/stave/pkg/sh"
)

// Lint runs markdownlint on all markdown files.
func Lint() error {
	return sh.RunV("markdownlint", ".")
}

