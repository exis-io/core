package main

import "fmt"
import riffle "github.com/exis-io/core/riffle"

func main() {
	// set flags for testing
	riffle.SetFabricLocal()
	riffle.SetLogLevelDebug()

	// Domain objects
	app := riffle.NewDomain("xs.damouse")
	receiver := app.Subdomain("receiver")

	// Connect
	receiver.Join()

	receiver.Subscribe("sub1", func(a string) {
		fmt.Println("sub1 Got string ", a)
	})
	receiver.Subscribe("sub2", func(a int64) {
		fmt.Println("sub2 Got int ", a)
	})
	receiver.Subscribe("sub3", func(a int64) {
		fmt.Println("sub3 Got int ", a)
	})

	receiver.Register("reg1", func(a string) {
		fmt.Println("reg1 Got string ", a)
	})
	receiver.Register("reg2", func(a int64) {
		fmt.Println("reg2 Got int ", a)
	})
	receiver.Register("reg3", func(a int64) {
		fmt.Println("reg3 Got int ", a)
	})

	// Handle until the connection closes
	receiver.Listen()
}