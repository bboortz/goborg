package borg

import (
	"time"
)

type Borg struct {
	Id       string    `json:"id"`
	Addr     string    `json:"addr"`
	LastSeen time.Time `json:"lastseen"`
}

type Borgs []Borg

func NewBorg(id string, addr string) Borg {
	b := Borg{
		Id:       id,
		Addr:     addr,
		LastSeen: time.Now(),
	}
	return b
}
