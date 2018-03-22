package federation

import (
	"wubba/lubba/apub"
	"wubba/lubba/pubsub"
)

func NewWebSubTransport() Transport {
	return &pubsub.Transport{}
}

type Federator struct {
	transports []Transport
}

func (f *Federator) Broadcast(post apub.Post) error {
	for idx := range f.transports {
		go f.transports[idx].Broadcast(post)
	}
	return nil
}
