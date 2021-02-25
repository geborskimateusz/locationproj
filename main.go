package main

import (
	"github.com/gin-gonic/gin"
)

const (
	nearbyUsersRoute   string = "api/v1/users/findNearby"
	matchesStatusRoute string = "api/v1/postinvitation/matchesStatus/:currentUser"
)

func main() {
	InitializeUsers()
	InitializeMatches()
	ServerInstance().Run()
}

func ServerInstance() *gin.Engine {

	router := gin.Default()

	router.GET(nearbyUsersRoute, NearbyUsers)
	router.GET(matchesStatusRoute, MatchesStatus)

	go h.run()

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	return router
}
