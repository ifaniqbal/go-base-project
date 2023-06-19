package main

import (
	"fmt"
	"github.com/ifaniqbal/go-base-project/internal/api"
)

func main() {
	app := api.Default()
	if err := app.Start(); err != nil {
		panic(fmt.Errorf("app.Start: %w", err))
	}
}
