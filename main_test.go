package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) {
	TestingT(t)
}

type MainSuite struct{}

var _ = Suite(&MainSuite{})

func (s *MainSuite) TestHello(c *C) {
	c.Assert(hello(), Equals, "Hello World!")
}
