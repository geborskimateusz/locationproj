package main

import (
	"github.com/gin-gonic/gin"
)

func NearbyUsers(c *gin.Context) {
	c.JSON(200, gin.H{"users": GetUsersNearby()})
}

func MatchesStatus(c *gin.Context) {
	currentUser := c.Param("currentUser")

	c.JSON(200, GetMatches(currentUser))
}
