package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	d := time.Date(2025, time.April, 7, 14, 15, 0, 0, time.Local)
	dun := time.Until(d)

	return int64(dun.Hours()) / 24
}
