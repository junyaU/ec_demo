package main

import "github.com/demo/layer/infra"

func main() {
	bus := infra.InitBus()
	infra.InitRouter(bus)
}
