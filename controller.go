package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func NearbyUsers(c *gin.Context) {

	firstUser := User{
		Nickname: "superuser123",
		Coordinates: Coordinates{
			X: 50.868529712179935,
			Y: 19.109939369334157,
		},
	}

	secondUser := User{
		Nickname: "anonuser",
		Coordinates: Coordinates{
			X: 50.866151341448436,
			Y: 19.11058980503063,
		},
	}

	thirdUser := User{
		Nickname: "thirdUSer",
		Coordinates: Coordinates{
			X: 50.856700985843815,
			Y: 19.108540597332587,
		},
	}
	fourthUser := User{
		Nickname: "JohnyFourthUserDoe",
		Coordinates: Coordinates{
			X: 50.8515960533141255,
			Y: 19.109704386867137,
		},
	}

	usersNearby := []User{}
	usersNearby = append(usersNearby, firstUser)
	usersNearby = append(usersNearby, secondUser)
	usersNearby = append(usersNearby, thirdUser)
	usersNearby = append(usersNearby, fourthUser)
	log.Println(usersNearby)

	c.JSON(200, gin.H{"users": usersNearby})
}
