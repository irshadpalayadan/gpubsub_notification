package main

import (
	"github.com/gin-gonic/gin"
	notify "github.com/irshadpalayadan/gpubsub_notification/module/notification"
)

func getServerStatus(ctx *gin.Context) {
	ctx.JSON(200, "notification server is healthy")
}

func main() {
	notify.InitializeNotification()
	router := gin.Default()

	router.GET("/status", getServerStatus)
	router.GET("/notification/:userid", notify.GetNotificationForUserId)
	router.POST("/notification/:userid/all", notify.MarkAllNotificationReadForUser)
	router.DELETE("/notification/:userid/all", notify.ClearAllNotificationForUser)

	router.Run(":8002")
}
