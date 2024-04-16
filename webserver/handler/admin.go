package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"lxtend.com/friend_link/webserver/service"
)

func DeleteFriendLink(c *gin.Context) {
	name := c.Query("name")
	if err := service.DeleteFriend(name); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true})
}
