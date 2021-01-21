package registry

// Tag represents a google artifact repository image tag
type Tag struct {
	Tag		string
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`
}

// Tags respresents the object returned from the tags endpoint
type Tags struct {
	Tags []Tag `json:"tags,omitempty"`
}
