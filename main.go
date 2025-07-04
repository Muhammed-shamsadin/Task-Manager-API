package main

import (
	"example.com/task_manager/data"
	"example.com/task_manager/router"
)

func main() {
	data.ConnectMongoDB()
	r := router.SetupRouter() // Assuming you have a function to set up your router
	r.Run(":8080") // Start the server on port 8080

}
