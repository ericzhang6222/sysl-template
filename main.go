// main.go
package main

import (
	"context"

	"github.com/anz-bank/sysl-template/internal/server"
)

func main() {
	// Now the LoadServices function is called to start our server
	server.LoadServices(context.Background())
}
