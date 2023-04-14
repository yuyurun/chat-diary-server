package main

import (
	"chat-diary-server/interface/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}
