package main

import "time"

//message represents a single message

type Message struct {
	Name    string
	Message string
	When    time.Time
	AvatarURL string
}