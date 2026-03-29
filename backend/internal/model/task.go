package model

import "time"

type Task struct {
	ID        int64
	Title     string `binding:"required"`
	Completed bool
	CreatedAt time.Time `binding:"required"`
}
