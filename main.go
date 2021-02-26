package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

const (
	nearbyUsersRoute   string = "api/v1/users/findNearby"
	matchesStatusRoute string = "api/v1/matchesStatus/:currentUser"
	inviteRoute        string = "api/v1/invite/:currentUser"
)

func main() {
	fmt.Printf("Initializing Users DB\n")
	InitializeUsers()
	fmt.Println(usersNearby)

	fmt.Printf("Initializing Matches DB\n")
	InitializeMatches()
	fmt.Println(matches)

	ServerInstance().Run()
}

func ServerInstance() *gin.Engine {

	router := gin.Default()

	router.GET(nearbyUsersRoute, NearbyUsers)
	router.GET(matchesStatusRoute, FindMatchesStatusByUsername)
	router.PUT(matchesStatusRoute, UpdateMatchesStatusByUsername)
	router.POST(inviteRoute, Invite)

	go h.run()

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	return router
}
