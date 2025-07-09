package main

import (
	"fmt"
	"zartool/internal/database"
)

func main() {
	server := database.InitServer()

	fmt.Println("server is running: ", server.Addr)
}
