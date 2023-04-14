package main

import (
	"chat-diary-server/internal/app/router"
)

func main() {
	r := router.SetupRouter()
	r.Run(":8080")
}