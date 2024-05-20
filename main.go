package main

var db = make(map[string]string)

func main() {
	r := makeRoutes()

	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
