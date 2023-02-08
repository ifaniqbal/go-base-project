package main

import (
	"fmt"

	"bitbucket.org/ifan-moladin/base-project/modules/consumer"
)

func main() {
	app := consumer.Default()
	if err := app.Run(); err != nil {
		panic(fmt.Errorf("app.Start: %w", err))
	}
}
