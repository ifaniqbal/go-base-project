package main

import (
	"fmt"
	"github.com/ifaniqbal/go-base-project/internal/consumer"
)

func main() {
	app := consumer.Default()
	if err := app.Run(); err != nil {
		panic(fmt.Errorf("app.Start: %w", err))
	}
}
