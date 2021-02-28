package main

func main() {
	InitializeUsers()
	InitializeMatches()

	//chat websocket hub
	go h.run()

	RouterInstance().Run()
}
