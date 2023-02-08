package main

import (
	"fmt"

	"bitbucket.org/ifan-moladin/base-project/modules/api"
)

func main() {
	app := api.Default()
	if err := app.Start(); err != nil {
		panic(fmt.Errorf("app.Start: %w", err))
	}
}
