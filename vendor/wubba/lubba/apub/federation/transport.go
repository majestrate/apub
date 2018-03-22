package federation

import "wubba/lubba/apub"

// Transport defines a one way post federation transport
type Transport interface {
	// Broadcast sends a notifcation to all subscribed users on this transport
	Broadcast(post apub.Post) error
}
