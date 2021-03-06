package main

import "github.com/gin-gonic/gin"

const (
	nearbyUsersRoute   string = "/users/findNearby"
	matchesStatusRoute string = "/matchesStatus/:currentUser"
	inviteRoute        string = "/invite/:currentUser"
	chatRoute          string = "/ws/:currentUser/connectedTo/:connectedTo"
)

func RouterInstance() *gin.Engine {

	router := gin.Default()
	router.Group("api/v1")

	router.GET(nearbyUsersRoute, NearbyUsers)
	router.GET(matchesStatusRoute, FindMatchesStatusByUsername)
	router.PUT(matchesStatusRoute, UpdateMatchesStatusByUsername)
	router.POST(inviteRoute, Invite)
	router.GET(chatRoute, JoinChatroom)

	return router
}
