package main

import "go-pipeliner/src/app"

func main() {
	app.NewApp().Run().WithGracefulShutdown()
}
