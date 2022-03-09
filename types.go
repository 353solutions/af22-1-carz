package carz

import "time"

type Location struct {
	Lat float64
	Lng float64
}

type Status byte

const (
	StatusOff = iota
	StatusFree
	StatusRide
)

type UpdateMessage struct {
	CarID    string
	DriverID string
	Location Location
	Time     time.Time
	Status   Status
}
