package config

import "time"

const (
	SendEmail     = false
	EmailFrom     = ""
	EmailTestTo   = ""
	EmailHost     = ""
	EmailUsername = ""
	EmailPassword = ""
	EmailPort     = 465
	// EmailTimeout  = 80 * time.Millisecond
	EmailTimeout = 10 * time.Second
)
