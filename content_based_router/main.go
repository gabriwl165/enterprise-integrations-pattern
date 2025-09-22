package content_based_router

import "strings"

type ContentBasedRouter struct {
	WidgetQueue chan string
	GadgetQueue chan string
	DunnoQueue  chan string
}

func (c *ContentBasedRouter) OnMessage(message string) {
	if strings.HasPrefix(message, "W") {
		c.WidgetQueue <- message
	} else if strings.HasPrefix(message, "G") {
		c.GadgetQueue <- message
	}
}

func NewContentBasedRouter() *ContentBasedRouter {
	return &ContentBasedRouter{
		WidgetQueue: make(chan string),
		GadgetQueue: make(chan string),
		DunnoQueue:  make(chan string),
	}
}
