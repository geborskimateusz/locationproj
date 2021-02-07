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

	router.GET(nearbyUsersRoute, NearbyUsers)

	return router
}
