package main

import (
	"azri_hub/router"
)

func main() {

	r := router.NewRouter()
	r.Run(":8081")
}
