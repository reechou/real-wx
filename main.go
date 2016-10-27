package main

import (
	"github.com/reechou/real-wx/config"
	"github.com/reechou/real-wx/controller"
)

func main() {
	controller.NewWXLogic(config.NewConfig()).Run()
}
