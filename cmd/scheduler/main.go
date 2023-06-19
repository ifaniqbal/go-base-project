package main

import (
	"fmt"
	"github.com/ifaniqbal/go-base-project/internal/scheduler"
)

func main() {
	app := scheduler.Default()
	if err := app.Run(); err != nil {
		panic(fmt.Errorf("app.Start: %w", err))
	}
}
