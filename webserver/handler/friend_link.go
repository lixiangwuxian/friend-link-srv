package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	emailhandler "lxtend.com/friend_link/email_handler"
	"lxtend.com/friend_link/logger"
	db_model "lxtend.com/friend_link/webserver/model"
	"lxtend.com/friend_link/webserver/service"
)

// ApproveFriendLink 批准友链申请
func ApproveFriendLink(c *gin.Context) {
	token := c.Query("token")
	data, _ := service.GetApplicationByApproveToken(token)
	if data.Email != "" {
		go emailhandler.GetMailHandler().SendApplicationApproved(data.Email)
	}
	c.Status(http.StatusOK)
}

// ApplicateFriendLink 申请友链
func ApplicateFriendLink(c *gin.Context) {
	var friendData db_model.ApplicationData
	if err := c.ShouldBindJSON(&friendData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	changeToken, approveToken, err := service.NewApplication(friendData.Name, friendData.Url, friendData.Description, friendData.Avatar, friendData.Email)
	if changeToken == "" && err == nil {
		logger.Info("Application already exists, User is ", friendData.Name)
		c.JSON(http.StatusOK, fmt.Sprintf("{\"token\":\"%s\"}", ""))
		return
	}
	if err != nil {
		logger.Error("Error creating application: ", err.Error())
		// w.WriteHeader(http.StatusInternalServerError)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	go emailhandler.GetMailHandler().SendFriendApplicationByEmail(friendData.Name, friendData.Url, friendData.Description, friendData.Avatar, approveToken)
	if friendData.Email != "" {
		go emailhandler.GetMailHandler().SendApplicationUploaded(friendData.Email)
	}
	c.JSON(http.StatusOK, fmt.Sprintf("{\"token\":\"%s\"}", changeToken))
}

// RejectFriendLink 拒绝友链申请
func RejectFriendLink(c *gin.Context) {
	token := c.Query("token")
	data, _ := service.GetApplicationByApproveToken(token)
	logger.Info(data.Name, " was rejected")
	go emailhandler.GetMailHandler().SendApplicationRejected(data.Email)
	service.DeleteApplication(token)
	c.JSON(http.StatusOK, gin.H{"success": true})
}

// GetFriendList 获取友链列表
func GetFriendList(c *gin.Context) {

	friendList, err := service.GetFriends()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, friendList)
}
