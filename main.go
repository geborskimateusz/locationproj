package main

import (
	"github.com/gin-gonic/gin"
)

const (
	nearbyUsersRoute string = "api/users/findNearby"
)

func main() {
	ServerInstance().Run()
}

func ServerInstance() *gin.Engine {

	router := gin.Default()
	// router.GET(nearbyUsersRoute, NearbyUsers)

	go h.run()

	router.LoadHTMLFiles("index.html")

	router.GET("/room/:roomId", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	router.GET("/ws/:roomId", func(c *gin.Context) {
		roomId := c.Param("roomId")
		serveWs(c.Writer, c.Request, roomId)
	})

	return router
}
