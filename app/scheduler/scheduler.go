package main

import (
	"fmt"

	"bitbucket.org/ifan-moladin/base-project/modules/scheduler"
)

func main() {
	app := scheduler.Default()
	if err := app.Run(); err != nil {
		panic(fmt.Errorf("app.Start: %w", err))
	}
}
