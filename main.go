package main

import "router"
import "fmt"

const ListenAddress = ":8090"

func main() {
	router.Start(ListenAddress)
	fmt.Printf("Server listening on: %s ...", ListenAddress)
	fmt.Println()
	// List to quit channel here later, quit gracefully
	select {}

}
