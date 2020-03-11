package main

import "github.com/panupong25509/go-structure/routes"

func main() {
	e := routes.Init()
	e.Logger.Fatal(e.Start(":8080"))
}
