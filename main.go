package main

import (
	"github.com/babelcoder-enterprise-courses/go-fiber-testing/server"
)

func main() {
	s := server.New()
	s.Setup()
	s.Start()
}
