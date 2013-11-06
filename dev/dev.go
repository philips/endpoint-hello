package main

import (
	"code.google.com/p/rsc/devweb/slave"
	_ "github.com/philips/endpoint-hello/service"
	_ "github.com/philips/endpoint-hello/api"
)

func main() {
	slave.Main()
}
