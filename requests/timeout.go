package requests

import "time"

var requestTimeout = time.Duration(3000) *time.Millisecond

func SetTimeout(timeout uint64) {
	requestTimeout = time.Duration(timeout) * time.Millisecond
}
