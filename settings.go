package main

import "time"

type Settings struct {
	soundEnabled       bool
	autoStartEnabled   bool
	notificationVolume float64
	timer              time.Duration
}
