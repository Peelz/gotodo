package models

import (
	"errors"
	"fmt"
	"time"
)

type TaskStatus string

const (
	Pending TaskStatus = "pending"
	Success TaskStatus = "success"
	Fail    TaskStatus = "fail"
	Cancel  TaskStatus = "cancel"
)
// Task model
type Task struct {
	ID        uint
	Name   string
	Status TaskStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Get value from TaskStatus
func (t TaskStatus) Value() (string, error) {
	switch t {
	case Pending, Success, Fail, Cancel: //valid case
		return string(t), nil
	}
	return "", errors.New("invalid product type value") //else is invalid
}

// Scan value is valid for TaskStatus
func (t *TaskStatus) Scan(value interface{}) error {
	var pt TaskStatus
	if value == nil {
		*t = ""
		return nil
	}
	st, ok := value.(string)
	if !ok {
		return errors.New("invalid data for task status")
	}
	pt = TaskStatus(st) //convert type from string to TaskStatus
	switch pt {
	case Pending, Success, Fail, Cancel: //valid case
		*t = pt
		return nil
	}
	return fmt.Errorf("invalid product type value :%s", st) //else is invalid
}
