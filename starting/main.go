package main

import (
	"rsc.io/quote"
	"github.com/danny/gomods/testing"
	"github.com/danny/gomods/times"
)

func main() {
	println(quote.Hello())
	testing.HelloThere()
	times.TimeNow()

	//note: Folder name must be the same as package name for imports
}
