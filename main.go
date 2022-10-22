package main

import (
	"MyToDoList/conf"
	"MyToDoList/routes"
)

func main() {
	conf.Init()

	r := routes.NewRouter()
	_ = r.Run(conf.HttpPort)
}
