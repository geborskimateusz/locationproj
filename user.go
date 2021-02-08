package main

type User struct {
	Nickname    string      `json:"nickname"`
	Coordinates Coordinates `json:"coordinates"`
}

type Coordinates struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}
