package main

type User struct {
	Nickname    string      `json:"nickname"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	X string `json:"x"`
	Y string `json:"y"`
}
