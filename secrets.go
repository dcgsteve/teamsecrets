// manage secrets

package main

import (
	"errors"
	"fmt"
)

func (c *GetCmd) Run(ctx *Context) error {
	if ctx.CurrentProject == "" {
		return errors.New("project is not set")
	}

	// TODO: get secret
	s := "dummy"

	fmt.Printf("%s", s)
	return nil
}

func (c *SetCmd) Run(ctx *Context) error {
	if ctx.CurrentProject == "" {
		return errors.New("project is not set")
	}

	// TODO: store secret

	return nil
}
