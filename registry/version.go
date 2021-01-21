package registry

import (
	"time"
)

// Version represents a google artifact repository image version
type Version struct {
	Digest	   string
	Name       string    `json:"name,omitempty"`
	CreateTime time.Time `json:"createTime,omitempty"`
	UpdateTime time.Time `json:"updateTime,omitempty"`
}

// Versions respresents the object returned from the versions endpoint
type Versions struct {
	Versions []Version `json:"versions,omitempty"`
}
