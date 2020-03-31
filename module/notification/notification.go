package notification

import (
	"fmt"

	"github.com/gin-gonic/gin"
	gpubsub "github.com/irshadpalayadan/gpubsub_notification/module/pubsub"
)

type Notify struct {
	id      string
	Message string `json:"message"`
	UserId  string `json:"userid"`
	Watched bool   `json:"watched"`
}

var NotifyCache []Notify

func InitializeNotification() {
	NotifyCache = append(NotifyCache, Notify{"120", "hello world", "1", false})
}

func GetNotificationForUserId(ctx *gin.Context) {

	td := gpubsub.SubscribePull("gpubsub-272721", "signup_notify")
	fmt.Print(td)

	userId := ctx.Param("userid")
	if userId == "" {
		ctx.JSON(400, gin.H{"status": "failed", "msg": "invalid payload"})
		return
	}

	var notifications = []Notify{}

	for _, item := range NotifyCache {

		if item.UserId == userId {
			notifications = append(notifications, item)
		}
	}

	ctx.JSON(200, gin.H{"status": "success", "data": notifications})
}

func MarkAllNotificationReadForUser(ctx *gin.Context) {

}

func ClearAllNotificationForUser(ctx *gin.Context) {

}
