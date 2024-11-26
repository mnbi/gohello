// gohello/internal/greeting/greeting.go

package greeting

import (
	"math/rand"
)

const (
	HI = iota
	HELLO
	GREETING
)

type Message string

var messages = [...]Message{
	"Hi!",
	"Hello!",
	"Hello, world!",
}

func Greeting() Message {
	return messages[rand.Intn(len(messages))]
}
