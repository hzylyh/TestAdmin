package main

func main() {
	router := MapRoutes()
	router.Run(":8089")
}
