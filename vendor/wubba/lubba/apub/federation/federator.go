package federation

import "wubba/lubba/apub"

type Federator interface {
	Send(apub.Post)
}
