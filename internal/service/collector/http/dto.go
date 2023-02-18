package http

import "time"

type DTO struct {
	Program     string        `json:"program"`
	DurationMS  time.Duration `json:"durationMS"`
	PathProject string        `json:"pathProject"`
}
