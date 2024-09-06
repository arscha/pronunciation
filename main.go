package main

import (
	"github.com/arscha/pronunciation/api"
)

func main() {
	api.HandleRoutes(":9000")
}
