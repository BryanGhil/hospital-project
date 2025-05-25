package main

import "backend/app"

func main() {
	a := app.App{}

	a.Init()
	a.Run()
}