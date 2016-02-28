package main

import (
	"github.com/xh3b4sd/anna/spec"
)

func (a *anna) MarshalJSON() ([]byte, error) {
	a.Log.WithTags(spec.Tags{L: "D", O: a, T: nil, V: 15}, "call MarshalJSON")
	return nil, nil
}