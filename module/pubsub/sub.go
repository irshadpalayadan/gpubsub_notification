package gpubsub

import (
	"context"
	"fmt"
	"sync"

	"cloud.google.com/go/pubsub"
	"github.com/sirupsen/logrus"
)

type publishReturn struct {
	status string
	msg    string
	data   string
}

func SubscribePull(projectID, subscriberID string) *publishReturn {

	ctx := context.Background()

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		logrus.Fatal(err)
		return &publishReturn{status: "failure", msg: "google pubsup connection failed"}
	}

	sub := client.Subscription(subscriberID)
	cctx, cancel := context.WithCancel(ctx)

	var message string
	var mu sync.Mutex
	err = sub.Receive(cctx, func(_ctx context.Context, msg *pubsub.Message) {

		message = string(msg.Data)
		msg.Ack()
		mu.Lock()
		defer mu.Unlock()
		cancel()
	})

	if err != nil {
		logrus.Fatal(err)
		return &publishReturn{status: "failure", msg: "google pubsup msg receive failed"}
	}

	fmt.Print("message : " + message)

	return &publishReturn{status: "success", msg: "google pubsup msg receive success", data: message}
}
