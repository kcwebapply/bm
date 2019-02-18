package main

import (
	"os"

	app "github.com/kcwebapply/bm/app"
)

func main() {
	app := app.InitApp()
	app.Run(os.Args)
}
