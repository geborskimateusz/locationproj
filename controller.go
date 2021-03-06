package main

import (
	"github.com/gin-gonic/gin"
)

func NearbyUsers(c *gin.Context) {
	c.JSON(200, gin.H{"users": GetUsersNearby()})
}

func FindMatchesStatusByUsername(c *gin.Context) {
	currentUser := c.Param("currentUser")
	c.JSON(200, GetMatches(currentUser))
}

func UpdateMatchesStatusByUsername(c *gin.Context) {
	currentUser := c.Param("currentUser")

	var matches Matches
	err := c.ShouldBindJSON(&matches)
	if err != nil {
		c.JSON(200, "Wrong JSON structure")
		return
	}

	c.JSON(200, UpdateMatches(currentUser, matches))
}

func Invite(c *gin.Context) {
	currentUser := c.Param("currentUser")

	invitedUser := struct {
		Username string
	}{}

	err := c.ShouldBindJSON(&invitedUser)
	if err != nil {
		c.JSON(200, "Wrong JSON strucure")
		return
	}

	c.JSON(200, MatchRequest(currentUser, invitedUser.Username))
}

func JoinChatroom(c *gin.Context) {
	currentUser := c.Param("currentUser")
	connectedTo := c.Param("connectedTo")

	if ok := CheckMatch(currentUser, connectedTo); ok {
		channel := GenerateUniqueChannel(currentUser, connectedTo)
		serveWs(c.Writer, c.Request, channel)
	} else {
		c.JSON(404, "Channel does not exist")
	}
}
