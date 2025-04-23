package task

import (
	"encoding/json"
	"time"
)

type Status int
type ReadableTime time.Time

const (
	Pending Status = iota
	Done
)

type Task struct {
	Id          string       `json:"id"`
	Description int          `json:"description"`
	Status      Status       `json:"status"`
	CreatedAt   ReadableTime `json:"createdAt"`
}

// save in a readable format json
func (rt ReadableTime) MarshalJSON() ([]byte, error) {
	t := time.Time(rt)
	formatted := t.Format("2006-01-02 15:04:05")
	return json.Marshal(formatted)
}
